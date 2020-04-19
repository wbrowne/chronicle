package lb

import (
	"context"
	"fmt"
	"log"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/attributes"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/serviceconfig"

	api "github.com/wbrowne/chronicle/api/v1"
)

type Resolver struct {
	mu sync.Mutex

	clientConn    resolver.ClientConn
	resolverConn  *grpc.ClientConn
	serviceConfig *serviceconfig.ParseResult
}

const Name = "chronicle"

var _ resolver.Builder = (*Resolver)(nil)

// setup client conn to server
func (r *Resolver) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	var dialOpts []grpc.DialOption
	if opts.DialCreds != nil {
		dialOpts = append(
			dialOpts,
			grpc.WithTransportCredentials(opts.DialCreds),
		)
	}

	r.clientConn = cc // should be above?

	r.serviceConfig = r.clientConn.ParseServiceConfig(
		fmt.Sprintf(`{"loadBalancingConfig":[{"%s":{}}]}`, Name),
	)

	var err error
	r.resolverConn, err = grpc.Dial(target.Endpoint, dialOpts...)
	if err != nil {
		return nil, err
	}
	r.ResolveNow(resolver.ResolveNowOptions{})

	return r, nil
}

// custom resolver scheme name (otherwise default is DNS)
func (r *Resolver) Scheme() string {
	return Name
}

func init() {
	resolver.Register(&Resolver{})
}

var _ resolver.Builder = (*Resolver)(nil)

// resolve targets and update client conn
func (r *Resolver) ResolveNow(resolver.ResolveNowOptions) {
	r.mu.Lock()
	defer r.mu.Unlock()

	client := api.NewLogClient(r.resolverConn)

	// get cluster and then set on cc attributes
	res, err := client.GetServers(context.Background(), &api.GetServersRequest{})
	if err != nil {
		log.Printf(
			"[ERROR] chronicle: failed to resolve servers: %v",
			err,
		)
		return
	}

	var addrs []resolver.Address
	for _, server := range res.Servers {
		addrs = append(addrs, resolver.Address{
			Addr: server.RpcAddr,
			Attributes: attributes.New(
				"is_leader",
				server.IsLeader,
			),
		})
	}

	r.clientConn.UpdateState(resolver.State{
		Addresses:     addrs,
		ServiceConfig: r.serviceConfig,
	})
}

func (r *Resolver) Close() {
	if err := r.resolverConn.Close(); err != nil {
		log.Printf("[ERROR] chronicle: failed to close conn: %v", err)
	}
}
