package static

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/upfluence/pkg/discovery/peer"
	"github.com/upfluence/pkg/discovery/resolver"
)

type Resolver struct {
	peers []*peer.Peer

	resolverChans []chan resolver.Update
	muChans       sync.Mutex
}

func NewResolverFromStrings(addrs []string) *Resolver {
	var peers = make([]*peer.Peer, len(addrs))

	for i, addr := range addrs {
		peers[i] = &peer.Peer{Addr: addr}
	}

	return &Resolver{peers: peers}
}

func (r *Resolver) String() string {
	var addrs = make([]string, len(r.peers))

	for i, peer := range r.peers {
		addrs[i] = peer.Addr
	}

	return fmt.Sprintf("resolver/static: [seeds: %v]", addrs)
}

func (r *Resolver) Open(_ context.Context) error { return nil }

func (r *Resolver) Close() error {
	for _, ch := range r.resolverChans {
		close(ch)
	}

	return nil
}

func (r *Resolver) Resolve(_ context.Context) (<-chan resolver.Update, error) {
	var ch = make(chan resolver.Update, 0)

	go func() { time.Sleep(5 * time.Second); ch <- resolver.Update{Additions: r.peers} }()

	r.muChans.Lock()
	r.resolverChans = append(r.resolverChans, ch)
	r.muChans.Unlock()

	return ch, nil
}