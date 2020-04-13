package server

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"

	api "github.com/wbrowne/chronicle/api/v1"
	"github.com/wbrowne/chronicle/internal/auth"
	sec "github.com/wbrowne/chronicle/internal/conf"
	"github.com/wbrowne/chronicle/internal/log"
)

func TestServer(t *testing.T) {
	for scenario, fn := range map[string]func(
		t *testing.T,
		rootClient, nobodyClient api.LogClient,
		config *Config,
	){
		"produce/consume a message to/from the log succeeds": testProduceConsume,
		"produce/consume stream succeeds":                    testProduceConsumeStream,
		"consume past log boundary fails":                    testConsumePastBoundary,
		"unauthorized client produce fails":                  testUnauthorized,
	} {
		t.Run(scenario, func(t *testing.T) {
			rootClient, nobodyClient, config, teardown := testSetup(t, nil)
			defer teardown()
			fn(t, rootClient, nobodyClient, config)
		})
	}
}

func testSetup(t *testing.T, fn func(*Config)) (rootClient, nobodyClient api.LogClient, config *Config, teardown func()) {
	t.Helper()

	network, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)

	// setup server

	dir, err := ioutil.TempDir("", "server_test")
	require.NoError(t, err)
	defer os.RemoveAll(dir)

	clog, err := log.NewLog(dir, &log.Config{})
	require.NoError(t, err)

	authorizer := auth.New(sec.ACLModelFile, sec.ACLPolicyFile)
	require.NoError(t, err)

	config = &Config{
		CommitLog:  clog,
		Authorizer: authorizer,
	}
	if fn != nil {
		fn(config)
	}

	serverTLSConfig, err := sec.SetupTLSConfig(sec.TLSConfig{
		CertFile:      sec.ServerCertFile,
		KeyFile:       sec.ServerKeyFile,
		CAFile:        sec.CAFile,
		ServerAddress: network.Addr().String(),
		Server:        true,
	})
	require.NoError(t, err)

	serverCreds := credentials.NewTLS(serverTLSConfig)
	server, err := NewGRPCServer(config, grpc.Creds(serverCreds))
	require.NoError(t, err)

	// goroutine since .Serve is blocking
	go func() {
		if err := server.Serve(network); err != nil {
			fmt.Errorf("failed to serve: %s", err)
		}
	}()

	// setup clients
	rootConn, rootClient := newClient(t, network.Addr().String(), sec.RootClientCertFile, sec.RootClientKeyFile)

	nobodyConn, nobodyClient := newClient(t, network.Addr().String(), sec.NobodyClientCertFile, sec.NobodyClientKeyFile)

	return rootClient, nobodyClient, config, func() {
		server.Stop()
		rootConn.Close()
		nobodyConn.Close()
		network.Close()
	}
}

func testProduceConsume(t *testing.T, client, _ api.LogClient, config *Config) {
	ctx := context.Background()
	want := &api.Record{
		Value: []byte("hello world"),
	}

	produce, err := client.Produce(
		context.Background(),
		&api.ProduceRequest{
			Record: want,
		},
	)
	require.NoError(t, err)

	consume, err := client.Consume(ctx, &api.ConsumeRequest{
		Offset: produce.Offset,
	})
	require.NoError(t, err)
	require.Equal(t, want.Value, consume.Record.Value)
	require.Equal(t, want.Offset, consume.Record.Offset)
}

func testConsumePastBoundary(t *testing.T, client, _ api.LogClient, config *Config) {
	ctx := context.Background()
	produce, err := client.Produce(ctx, &api.ProduceRequest{
		Record: &api.Record{
			Value: []byte("hello world"),
		},
	})
	require.NoError(t, err)

	consume, err := client.Consume(ctx, &api.ConsumeRequest{
		Offset: produce.Offset + 1,
	})
	if consume != nil {
		t.Fatal("consume not nil")
	}
	got := status.Code(err)
	want := status.Code(api.ErrOffsetOutOfRange{}.GRPCStatus().Err())
	if got != want {
		t.Fatalf("got err: %v, want: %v", got, want)
	}
}

func testProduceConsumeStream(t *testing.T, client, _ api.LogClient, config *Config) {
	ctx := context.Background()
	records := []*api.Record{
		{
			Value: []byte("first message"),
		},
		{
			Value: []byte("second message"),
		}}

	{
		stream, err := client.ProduceStream(ctx)
		require.NoError(t, err)

		for offset, record := range records {
			err = stream.Send(&api.ProduceRequest{
				Record: record,
			})
			require.NoError(t, err)

			res, err := stream.Recv()
			require.NoError(t, err)

			if res.Offset != uint64(offset) {
				t.Fatalf(
					"got offset: %d, want: %d",
					res.Offset,
					offset,
				)
			}
		}
	}
	{
		stream, err := client.ConsumeStream(
			ctx,
			&api.ConsumeRequest{Offset: 0},
		)
		require.NoError(t, err)
		for _, record := range records {
			res, err := stream.Recv()
			require.NoError(t, err)
			require.Equal(t, res.Record.Value, record.Value)
			require.Equal(t, res.Record.Offset, record.Offset)
		}
	}
}

func testUnauthorized(t *testing.T, _, client api.LogClient, config *Config, ) {
	ctx := context.Background()
	produce, err := client.Produce(context.Background(),
		&api.ProduceRequest{
			Record: &api.Record{
				Value: []byte("hello world"),
			},
		},
	)
	if produce != nil {
		t.Fatalf("produce response should be nil")
	}
	gotCode, wantCode := status.Code(err), codes.PermissionDenied
	if gotCode != wantCode {
		t.Fatalf("got code: %d, want: %d", gotCode, wantCode)
	}

	consume, err := client.Consume(ctx, &api.ConsumeRequest{
		Offset: 0,
	})
	if consume != nil {
		t.Fatalf("consume response should be nil")
	}
	gotCode, wantCode = status.Code(err), codes.PermissionDenied
	if gotCode != wantCode {
		t.Fatalf("got code: %d, want: %d", gotCode, wantCode)
	}
}

func newClient(t *testing.T, srvAdr, crtPath, keyPath string) (
	*grpc.ClientConn,
	api.LogClient,
) {
	tlsConfig, err := sec.SetupTLSConfig(sec.TLSConfig{
		CertFile: crtPath,
		KeyFile:  keyPath,
		CAFile:   sec.CAFile,
		Server:   false,
	})
	require.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tlsCreds := credentials.NewTLS(tlsConfig)
	conn, err := grpc.DialContext(ctx, srvAdr, grpc.WithTransportCredentials(tlsCreds), grpc.WithBlock())
	require.NoError(t, err)

	client := api.NewLogClient(conn)
	return conn, client
}
