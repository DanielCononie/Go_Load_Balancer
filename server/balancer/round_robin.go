package balancer

import (
	"net/url"
	"sync/atomic"
)

type RoundRobin struct {
	servers []*url.URL
	counter uint64
}

// Returns a new round robin instance with a count of 0 and servers passed from main.go
func NewRoundRobin(servers []*url.URL) *RoundRobin {
	return &RoundRobin{servers: servers}
}

// Determines which server to send the next request to, takes in the servers slice and counter and returns the url to send it to next
// by using %, if it is 1, it will use the first server, if it is on the last server already, it will start over at 0
func NextServer(rr *RoundRobin) *url.URL {
	index := atomic.AddUint64(&rr.counter, 1)
	return rr.servers[index%uint64(len(rr.servers))]
}
