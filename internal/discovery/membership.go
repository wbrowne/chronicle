package discovery

import (
	"log"
	"net"

	"github.com/hashicorp/serf/serf"
)

type Membership struct {
	Config
	handler Handler
	serf    *serf.Serf
	events  chan serf.Event
}

type Config struct {
	NodeName       string       // id
	BindAddr       *net.TCPAddr // gossip addr
	Tags           map[string]string
	StartJoinAddrs []string // address of nodes in cluster
}

type Handler interface {
	Join(name, addr string) error

	Leave(name string) error
}

func New(handler Handler, config Config) (*Membership, error) {
	m := &Membership{
		Config:  config,
		handler: handler,
	}
	if err := m.setupSerf(); err != nil {
		return nil, err
	}
	return m, nil
}

func (m *Membership) setupSerf() (err error) {
	config := serf.DefaultConfig()
	config.Init()

	config.MemberlistConfig.BindAddr = m.BindAddr.IP.String()
	config.MemberlistConfig.BindPort = m.BindAddr.Port
	config.MemberlistConfig.SecretKey = nil
	m.events = make(chan serf.Event)
	config.EventCh = m.events
	config.Tags = m.Tags
	config.NodeName = m.Config.NodeName
	m.serf, err = serf.Create(config)

	if err != nil {
		return err
	}

	go m.eventHandler() // start serf event handler

	// join the cluster
	if m.StartJoinAddrs != nil {
		_, err = m.serf.Join(m.StartJoinAddrs, true)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *Membership) eventHandler() {
	for e := range m.events {
		switch e.EventType() {
		case serf.EventMemberJoin:
			for _, member := range e.(serf.MemberEvent).Members {
				if m.isLocal(member) { // ignore self
					continue
				}
				m.handleJoin(member)
			}
		case serf.EventMemberLeave, serf.EventMemberFailed:
			for _, member := range e.(serf.MemberEvent).Members {
				if m.isLocal(member) { // ignore self
					return
				}
				m.handleLeave(member)
			}
		}
	}
}
func (m *Membership) handleJoin(member serf.Member) {
	if err := m.handler.Join(
		member.Name,
		member.Tags["rpc_addr"],
	); err != nil {
		log.Printf(
			"[ERROR] proglog: failed to join: %s %s",
			member.Name,
			member.Tags["rpc_addr"],
		)
	}
}

func (m *Membership) handleLeave(member serf.Member) {
	if err := m.handler.Leave(
		member.Name,
	); err != nil {
		log.Printf(
			"[ERROR] proglog: failed to leave: %s",
			member.Name,
		)
	}
}

// snapshot of current cluster members
func (m *Membership) Members() []serf.Member {
	return m.serf.Members()
}

func (m *Membership) isLocal(member serf.Member) bool {
	return m.serf.LocalMember().Name == member.Name
}

func (m *Membership) Leave() error {
	return m.serf.Leave()
}
