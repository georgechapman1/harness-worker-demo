package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"time"
)

type HealthResponse struct {
	Status    string    `json:"status"`
	Service   string    `json:"service"`
	Version   string    `json:"version"`
	Timestamp time.Time `json:"timestamp"`
}

type MetricsResponse struct {
	Goroutines int    `json:"goroutines"`
	GOOS       string `json:"goos"`
	GOARCH     string `json:"goarch"`
	Uptime     string `json:"uptime"`
}

var startTime = time.Now()

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

	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(MetricsResponse{
			Goroutines: runtime.NumGoroutine(),
			GOOS:       runtime.GOOS,
			GOARCH:     runtime.GOARCH,
			Uptime:     time.Since(startTime).String(),
		})
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Harness Worker Agent Service v%s\n", version)
	})

	fmt.Printf("[worker] Listening on :%s  version=%s\n", port, version)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		os.Exit(1)
	}
}
