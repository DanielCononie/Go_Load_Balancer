package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Response from Server 2 on port 8082\n")
	})
	fmt.Println("Server 2 running on :8082")
	http.ListenAndServe(":8082", nil)
}
