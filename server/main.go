package main

import (
	"fmt"
	"load_balancer/balancer"
	"load_balancer/metrics"
	"load_balancer/proxy"
	"log"
	"net/http"
	"net/url"
)

// CORS middleware
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	// Define backend servers
	servers := []*url.URL{
		{Scheme: "http", Host: "localhost:8081"},
		{Scheme: "http", Host: "localhost:8082"},
	}

	// Initialize load balancer
	load_balancer := balancer.NewRoundRobin(servers)

	// Handling forwarded requests
	http.Handle("/", enableCORS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Determines which server to send request to next
		target := balancer.NextServer(load_balancer)
		serverHost := target.Host

		// Metrics update for request start
		metrics.UpdateMetrics(serverHost)

		// Creates a new proxy to forward the request
		proxy := proxy.NewProxy(target)
		fmt.Printf("Forwarding request to: %s\n", target.Host)

		// Forwards the request and handles response
		proxy.ServeHTTP(w, r)

		// Metrics update for request completion
		metrics.CompleteRequest(serverHost)
	})))

	// Expose metrics endpoint
	http.Handle("/metrics", enableCORS(http.HandlerFunc(metrics.MetricsHandler)))

	log.Println("Load Balancer started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
