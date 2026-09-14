# Stage 1 — build
FROM golang:1.21-alpine AS builder
WORKDIR /build
COPY go.mod ./
COPY main.go ./
RUN go build -ldflags="-s -w" -o worker .

# Stage 2 — minimal runtime
FROM alpine:3.19
RUN apk --no-cache add ca-certificates wget \
 && addgroup -S worker \
 && adduser  -S worker -G worker
WORKDIR /app
COPY --from=builder /build/worker .
RUN chown -R worker:worker /app
USER worker
EXPOSE 8080
ENV PORT=8080
HEALTHCHECK --interval=15s --timeout=3s --retries=3 \
  CMD wget -qO- http://localhost:8080/healthz || exit 1
ENTRYPOINT ["./worker"]
