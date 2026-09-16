package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"sync/atomic"
	"time"
)

type HealthResponse struct {
	Status    string    `json:"status"`
	Service   string    `json:"service"`
	Version   string    `json:"version"`
	Timestamp time.Time `json:"timestamp"`
}

type ReadinessResponse struct {
	Ready    bool  `json:"ready"`
	Requests int64 `json:"requests_served"`
}

type MetricsResponse struct {
	Goroutines int    `json:"goroutines"`
	GOOS       string `json:"goos"`
	GOARCH     string `json:"goarch"`
	Uptime     string `json:"uptime"`
}

type VersionResponse struct {
	Version   string `json:"version"`
	GoVersion string `json:"go_version"`
	OS        string `json:"os"`
	Arch      string `json:"arch"`
}

var startTime = time.Now()
var requestCount int64

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
		atomic.AddInt64(&requestCount, 1)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(HealthResponse{
			Status:    "ok",
			Service:   "harness-worker",
			Version:   version,
			Timestamp: time.Now().UTC(),
		})
	})

	http.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt64(&requestCount, 1)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ReadinessResponse{
			Ready:    true,
			Requests: atomic.LoadInt64(&requestCount),
		})
	})

	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt64(&requestCount, 1)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(VersionResponse{
			Version:   version,
			GoVersion: runtime.Version(),
			OS:        runtime.GOOS,
			Arch:      runtime.GOARCH,
		})
	})

	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt64(&requestCount, 1)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(MetricsResponse{
			Goroutines: runtime.NumGoroutine(),
			GOOS:       runtime.GOOS,
			GOARCH:     runtime.GOARCH,
			Uptime:     time.Since(startTime).String(),
		})
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt64(&requestCount, 1)
		fmt.Fprintf(w, "Harness Worker Agent Service v%s\n", version)
	})

	fmt.Printf("[worker] Listening on :%s  version=%s\n", port, version)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		os.Exit(1)
	}
}
