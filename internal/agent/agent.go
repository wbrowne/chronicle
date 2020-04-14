package agent

import (
	"crypto/tls"
	"fmt"
	"net"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	api "github.com/wbrowne/chronicle/api/v1"
	"github.com/wbrowne/chronicle/internal/auth"
	"github.com/wbrowne/chronicle/internal/discovery"
	"github.com/wbrowne/chronicle/internal/log"
	"github.com/wbrowne/chronicle/internal/server"
)

type Agent struct {
	Config
	log          *log.Log
	server       *grpc.Server
	membership   *discovery.Membership
	replicator   *log.Replicator
	shutdown     bool
	shutdowns    chan struct{}
	shutdownLock sync.Mutex
}

type Config struct {
	ServerTLSConfig *tls.Config
	ClientTLSConfig *tls.Config
	DataDir         string
	BindAddr        *net.TCPAddr
	RPCPort         int
	NodeName        string
	StartJoinAddrs  []string
	ACLModelFile    string
	ACLPolicyFile   string
}

func (c Config) RPCAddr() string {
	return fmt.Sprintf("%s:%d", c.BindAddr.IP.String(), c.RPCPort)
}

func New(config Config) (*Agent, error) {
	a := &Agent{
		Config:    config,
		shutdowns: make(chan struct{}),
	}
	setup := []func() error{
		a.setupLog,
		a.setupServer,
		a.setupMembership,
	}
	for _, fn := range setup {
		if err := fn(); err != nil {
			return nil, err
		}
	}

	return a, nil
}

func (a *Agent) setupLog() error {
	var err error
	a.log, err = log.NewLog(
		a.Config.DataDir,
		&log.Config{},
	)
	return err
}

func (a *Agent) setupServer() error {
	serverConfig := &server.Config{
		CommitLog:  a.log,
		Authorizer: auth.New(
			a.Config.ACLModelFile,
			a.Config.ACLPolicyFile,
		),
	}
	var opts []grpc.ServerOption
	if a.Config.ServerTLSConfig != nil {
		creds := credentials.NewTLS(a.Config.ServerTLSConfig)
		opts = append(opts, grpc.Creds(creds))
	}
	var err error
	a.server, err = server.NewGRPCServer(serverConfig, opts...)
	if err != nil {
		return err
	}
	ln, err := net.Listen("tcp", a.Config.RPCAddr())
	if err != nil {
		return err
	}

	// start server
	go func() {
		if err := a.server.Serve(ln); err != nil {
			_ = a.Shutdown()
		}
	}()

	return err
}

func (a *Agent) setupMembership() error {
	var opts []grpc.DialOption
	if a.Config.ClientTLSConfig != nil {
		opts = append(opts,
			grpc.WithTransportCredentials(
				credentials.NewTLS(a.Config.ClientTLSConfig),
			),
		)
	}
	conn, err := grpc.Dial(a.Config.RPCAddr(), opts...)
	if err != nil {
		return err
	}
	client := api.NewLogClient(conn)
	a.replicator = &log.Replicator{
		DialOptions: opts,
		LocalServer: client,
	}
	a.membership, err = discovery.New(a.replicator, discovery.Config{
		NodeName: a.Config.NodeName,
		BindAddr: a.Config.BindAddr,
		Tags: map[string]string{
			"rpc_addr": a.Config.RPCAddr(),
		},
		StartJoinAddrs: a.Config.StartJoinAddrs,
	})

	return err
}

func (a *Agent) Shutdown() error {
	a.shutdownLock.Lock()
	defer a.shutdownLock.Unlock()

	if a.shutdown {
		return nil
	}
	a.shutdown = true

	close(a.shutdowns)
	shutdown := []func() error{
		a.membership.Leave,
		a.replicator.Close,
		func() error {
			a.server.GracefulStop()
			return nil
		},
		a.log.Close,
	}
	for _, fn := range shutdown {
		if err := fn(); err != nil {
			return err
		}
	}

	return nil
}
