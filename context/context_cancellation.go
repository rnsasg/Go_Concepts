package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

// Run Program : go run context_cancellation.go
// Open Browser : http://localhost:50001/
// Close the tab

// Output :
// Processing Request
// Request Cancelled

func main() {
	http.ListenAndServe(":50001", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		fmt.Fprint(os.Stdout, "Processing Request\n")
		select {
		case <-time.After(10 * time.Second):
			w.Write([]byte("Request Processed"))
		case <-ctx.Done():
			fmt.Fprint(os.Stderr, "Request Cancelledn")
		}
	}))
}
