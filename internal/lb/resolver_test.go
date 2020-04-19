package lb_test

import (
	"net"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/attributes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/serviceconfig"

	api "github.com/wbrowne/chronicle/api/v1"
	sec "github.com/wbrowne/chronicle/internal/conf"
	"github.com/wbrowne/chronicle/internal/lb"
	"github.com/wbrowne/chronicle/internal/server"
	"github.com/wbrowne/chronicle/internal/testutil"
)

func TestResolver(t *testing.T) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)

	tlsConfig, err := sec.SetupTLSConfig(sec.TLSConfig{
		CertFile:      testutil.ServerCertFile,
		KeyFile:       testutil.ServerKeyFile,
		CAFile:        testutil.CAFile,
		Server:        true,
		ServerAddress: "127.0.0.1",
	})
	require.NoError(t, err)

	serverCreds := credentials.NewTLS(tlsConfig)

	srv, err := server.NewGRPCServer(&server.Config{
		ServerInfo: &getServers{},
	}, grpc.Creds(serverCreds))
	require.NoError(t, err)

	go srv.Serve(l)

	tlsConfig, err = sec.SetupTLSConfig(sec.TLSConfig{
		CertFile:      testutil.RootClientCertFile,
		KeyFile:       testutil.RootClientKeyFile,
		CAFile:        testutil.CAFile,
		Server:        false,
		ServerAddress: "127.0.0.1",
	})
	require.NoError(t, err)

	clientCreds := credentials.NewTLS(tlsConfig)
	opts := resolver.BuildOptions{
		DialCreds: clientCreds,
	}

	conn := &clientConn{}

	r := &lb.Resolver{}
	_, err = r.Build(
		resolver.Target{
			Endpoint: l.Addr().String(),
		},
		conn,
		opts,
	)
	require.NoError(t, err)

	wantState := resolver.State{
		Addresses: []resolver.Address{{
			Addr:       "localhost:9001",
			Attributes: attributes.New("is_leader", true),
		}, {
			Addr:       "localhost:9002",
			Attributes: attributes.New("is_leader", false),
		}},
	}
	require.Equal(t, wantState, conn.state)

	conn.state.Addresses = nil
	r.ResolveNow(resolver.ResolveNowOptions{})
	require.Equal(t, wantState, conn.state)
}

type getServers struct{}

func (s *getServers) GetServers() ([]*api.Server, error) {
	return []*api.Server{{
		Id:       "leader",
		RpcAddr:  "localhost:9001",
		IsLeader: true,
	}, {
		Id:      "follower",
		RpcAddr: "localhost:9002",
	}}, nil
}

// mock for tracking resolver state
type clientConn struct {
	resolver.ClientConn
	state resolver.State
}

func (c *clientConn) UpdateState(state resolver.State) {
	c.state = state
}

func (c *clientConn) ReportError(err error) {}

func (c *clientConn) NewAddress(addrs []resolver.Address) {}

func (c *clientConn) NewServiceConfig(config string) {}

func (c *clientConn) ParseServiceConfig(config string, ) *serviceconfig.ParseResult {
	return nil
}
