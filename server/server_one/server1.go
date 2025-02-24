package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Response from Server 1 on port 8081\n")
	})
	fmt.Println("Server 1 running on :8081")
	http.ListenAndServe(":8081", nil)
}
