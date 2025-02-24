package health

import (
	"net/http"
	"time"
)

func IsServerAlive(url string) bool {
	client := http.Client{
		Timeout: 2 * time.Second,
	}
	resp, err := client.Get(url + "/health")
	if err != nil || resp.StatusCode != http.StatusOK {
		return false
	}
	return true
}
