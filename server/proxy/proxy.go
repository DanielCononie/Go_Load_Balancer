package proxy

import (
	"net/http/httputil"
	"net/url"
)

// Creates a new Reverse Proxy instance to handle modifying the request headers/url and when .ServeHTTP is called, it handles forwarding the request
func NewProxy(target *url.URL) *httputil.ReverseProxy {
	return httputil.NewSingleHostReverseProxy(target)
}
