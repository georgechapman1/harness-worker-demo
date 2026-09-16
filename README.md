# harness-worker-demo

A demo Go microservice showcasing Harness Worker Agents in a CI pipeline.

## Endpoints

| Endpoint | Description |
|----------|-------------|
| `GET /` | Service banner |
| `GET /healthz` | Health check (JSON) |
| `GET /metrics` | Runtime metrics (goroutines, GOOS, GOARCH, uptime) |

## Pipeline

Built, scanned, and published via **Harness Worker Agent - Build and Publish**:

1. **Build & Vet** — `go vet` + compile
2. **Smoke Test** — starts binary, hits `/healthz`
3. **Code Review** — AI-powered PR review (PR triggers only)
4. **Security Scan** — Semgrep, OWASP, OSV Scanner in parallel
5. **Publish** — Docker image → Harness Artifact Registry (`worker`)
