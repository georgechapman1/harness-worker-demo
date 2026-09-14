# harness-worker-demo

Dummy **Harness Worker Agent** microservice — a lightweight Go HTTP service
built and published to Harness Artifact Registry via a Harness CI pipeline.

## Endpoints

| Endpoint | Description |
|----------|-------------|
| `GET /` | Returns service name and version |
| `GET /healthz` | Returns JSON health status |

## Building locally

```bash
go build -o worker .
SERVICE_VERSION=local ./worker
curl http://localhost:8080/healthz
```

## Docker

```bash
docker build -t harness-worker-demo .
docker run -p 8080:8080 -e SERVICE_VERSION=local harness-worker-demo
```
