package metrics

import (
	"encoding/json"
	"net/http"
	"sync"
)

type Metrics struct {
	TotalRequests     uint64            `json:"total_requests"`
	RequestsPerServer map[string]uint64 `json:"requests_per_server"`
	ActiveConnections map[string]uint64 `json:"active_connections"`
	mu                sync.Mutex
}

var metrics = Metrics{
	RequestsPerServer: make(map[string]uint64),
	ActiveConnections: make(map[string]uint64),
}

// Update metrics when a request is sent
func UpdateMetrics(server string) {
	metrics.mu.Lock()
	defer metrics.mu.Unlock()
	metrics.TotalRequests++
	metrics.RequestsPerServer[server]++
	metrics.ActiveConnections[server]++
}

// Decrease active connections after the request finishes
func CompleteRequest(server string) {
	metrics.mu.Lock()
	defer metrics.mu.Unlock()
	metrics.ActiveConnections[server]--
}

// Expose the metrics as JSON
func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	metrics.mu.Lock()
	defer metrics.mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(metrics)
}
