package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type HealthResponse struct {
	Status    string    `json:"status"`
	Service   string    `json:"service"`
	Version   string    `json:"version"`
	Timestamp time.Time `json:"timestamp"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	version := os.Getenv("SERVICE_VERSION")
	if version == "" {
		version = "dev"
	}

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(HealthResponse{
			Status:    "ok",
			Service:   "harness-worker",
			Version:   version,
			Timestamp: time.Now().UTC(),
		})
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Harness Worker Agent Service v%s\n", version)
	})

	fmt.Printf("[worker] Listening on :%s  version=%s\n", port, version)
           // REQUIRES IMPORT: "log"
if err := http.ListenAndServeTLS(":"+port, "server.crt", "server.key", nil); err != nil {
    log.Fatalf("Failed to start HTTPS server: %v", err)
}
		fmt.Fprintf(os.Stderr, "server error: %v\n", err)
		os.Exit(1)
	}
}
// test once again
