# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN apk add --no-cache git
RUN go mod download

COPY . .
RUN go build -o server cmd/server/main.go

# Run stage
FROM alpine:latest

WORKDIR /app

# Install dependencies (if needed, e.g. for uploads)
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/server .
COPY --from=builder /app/.env.example .env

# Create uploads directory
RUN mkdir -p uploads/proofs

EXPOSE 8080

CMD ["./server"]
