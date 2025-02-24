package balancer

import (
	"net/url"
	"sync"
)

type Server struct {
	URL               *url.URL
	ActiveConnections int
}

type LeastConnections struct {
	servers []*Server
	mutex   sync.Mutex
}

func NewLeastConnections(servers []*url.URL) *LeastConnections {
	serverList := make([]*Server, len(servers))
	for i, server := range servers {
		serverList[i] = &Server{URL: server}
	}
	return &LeastConnections{servers: serverList}
}

func (lc *LeastConnections) NextServer() *Server {
	lc.mutex.Lock()
	defer lc.mutex.Unlock()

	var selected *Server
	for _, server := range lc.servers {
		if selected == nil || server.ActiveConnections < selected.ActiveConnections {
			selected = server
		}
	}

	selected.ActiveConnections++
	return selected
}

func (lc *LeastConnections) Release(server *Server) {
	lc.mutex.Lock()
	defer lc.mutex.Unlock()
	server.ActiveConnections--
}
