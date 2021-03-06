package agent_test

import (
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"testing"
	"time"

	"github.com/phayes/freeport"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"

	api "github.com/wbrowne/chronicle/api/v1"
	"github.com/wbrowne/chronicle/internal/agent"
	sec "github.com/wbrowne/chronicle/internal/conf"
	"github.com/wbrowne/chronicle/internal/lb"
)

func TestAgent(t *testing.T) {
	serverTLSConfig, err := sec.SetupTLSConfig(sec.TLSConfig{
		CertFile:      sec.ServerCertFile,
		KeyFile:       sec.ServerKeyFile,
		CAFile:        sec.CAFile,
		Server:        true,
		ServerAddress: "127.0.0.1",
	})
	require.NoError(t, err)

	clientTLSConfig, err := sec.SetupTLSConfig(sec.TLSConfig{
		CertFile:      sec.RootClientCertFile,
		KeyFile:       sec.RootClientKeyFile,
		CAFile:        sec.CAFile,
		Server:        false,
		ServerAddress: "127.0.0.1",
	})
	require.NoError(t, err)

	// setup 3 node cluster
	a1 := createAgent(t, serverTLSConfig, clientTLSConfig, "0", "")
	// two nodes join the first agent's cluster
	a2 := createAgent(t, serverTLSConfig, clientTLSConfig, "1", a1.BindAddr.String())
	a3 := createAgent(t, serverTLSConfig, clientTLSConfig, "2", a1.BindAddr.String())

	// verify agent shutdown
	defer func() {
		for _, a := range []*agent.Agent{a1, a2, a3} {
			err := a.Shutdown()
			require.NoError(t, err)
			require.NoError(t, os.RemoveAll(a.Config.DataDir))
		}
	}()

	// wait for agent discovery
	time.Sleep(3 * time.Second)

	leaderClient := client(t, a1, clientTLSConfig)
	recordVal := []byte("foo")

	produceResponse, err := leaderClient.Produce(
		context.Background(),
		&api.ProduceRequest{
			Record: &api.Record{
				Value: recordVal,
			},
		},
	)
	require.NoError(t, err)

	// wait until data replication has finished
	time.Sleep(3 * time.Second)

	consumeResponse, err := leaderClient.Consume(
		context.Background(),
		&api.ConsumeRequest{
			Offset: produceResponse.Offset,
		},
	)
	require.NoError(t, err)
	require.Equal(t, consumeResponse.Record.Value, recordVal)
	require.Equal(t, consumeResponse.Record.Offset, produceResponse.Offset)

	followerClient := client(t, a2, clientTLSConfig)
	consumeResponse, err = followerClient.Consume(
		context.Background(),
		&api.ConsumeRequest{
			Offset: produceResponse.Offset,
		},
	)
	require.NoError(t, err)
	require.Equal(t, consumeResponse.Record.Value, recordVal)
	require.Equal(t, consumeResponse.Record.Offset, produceResponse.Offset)

	// verify no more logs produced
	consumeResponse, err = leaderClient.Consume(
		context.Background(),
		&api.ConsumeRequest{
			Offset: produceResponse.Offset + 1,
		},
	)
	require.Nil(t, consumeResponse)
	require.Error(t, err)
	got := status.Code(err)
	want := status.Code(api.ErrOffsetOutOfRange{}.GRPCStatus().Err())
	require.Equal(t, got, want)
}

func createAgent(t *testing.T, serverTLSConfig, clientTLSConfig *tls.Config, id, agentAdr string) *agent.Agent {
	ports, err := freeport.GetFreePorts(2) // sd and rpc
	require.NoError(t, err)

	bindAddr := &net.TCPAddr{
		IP:   []byte{127, 0, 0, 1},
		Port: ports[0],
	}

	dataDir, err := ioutil.TempDir("", "agent_test")
	require.NoError(t, err)
	defer os.RemoveAll(dataDir)

	var startJoinAddrs []string
	if agentAdr != "" {
		startJoinAddrs = append(
			startJoinAddrs,
			agentAdr,
		)
	}
	a, err := agent.New(agent.Config{
		NodeName:        id,
		Bootstrap:       agentAdr == "",
		StartJoinAddrs:  startJoinAddrs,
		BindAddr:        bindAddr,
		RPCPort:         ports[1],
		DataDir:         dataDir,
		ACLModelFile:    sec.ACLModelFile,
		ACLPolicyFile:   sec.ACLPolicyFile,
		ServerTLSConfig: serverTLSConfig,
		ClientTLSConfig: clientTLSConfig,
	})
	require.NoError(t, err)

	return a
}

func client(t *testing.T, agent *agent.Agent, tlsConfig *tls.Config) api.LogClient {
	tlsCreds := credentials.NewTLS(tlsConfig)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(tlsCreds)}
	conn, err := grpc.Dial(fmt.Sprintf(
		"%s:///%s:%d",
		lb.Name,
		agent.Config.BindAddr.IP.String(),
		agent.Config.RPCPort,
	), opts...)
	require.NoError(t, err)

	client := api.NewLogClient(conn)

	return client
}
