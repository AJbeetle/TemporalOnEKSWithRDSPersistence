# ── Stage 1: Builder ──────────────────────────────────────────────────────────
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Copy dependency files first (better layer caching)
# Only re-downloads modules if go.mod or go.sum changed
COPY go.mod go.sum ./
RUN go mod download

# Copy entire source
COPY . .

# Build both binaries, output to /app/bin/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/bin/api     ./cmd/api
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/bin/worker  ./cmd/worker

# ── Stage 2: Runtime ──────────────────────────────────────────────────────────
FROM alpine:3.21

# ca-certificates needed for HTTPS calls (Stripe, RDS TLS, Temporal gRPC)
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

# Copy only the compiled binaries from builder — no Go toolchain in final image
COPY --from=builder /app/bin/api    ./api
COPY --from=builder /app/bin/worker ./worker

# Entrypoint script that reads the first argument
ENTRYPOINT ["sh", "-c", "\
  if [ \"$1\" = \"api\" ]; then \
    exec ./api; \
  elif [ \"$1\" = \"worker\" ]; then \
    exec ./worker; \
  else \
    echo \"Usage: docker run <image> [api|worker]\"; \
    exit 1; \
  fi", "--", "$@"]