# -----------------------------
# Build Stage
# -----------------------------
FROM golang:1.25-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build \
    -ldflags="-s -w" \
    -o api-gateway \
    ./cmd

# -----------------------------
# Runtime Stage
# -----------------------------
FROM alpine:3.22

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/api-gateway .

EXPOSE 8080

CMD ["./api-gateway"]
