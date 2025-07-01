# syntax=docker/dockerfile:1
FROM golang:1.24.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ecommerce-app ./cmd/main.go

# ---

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/ecommerce-app .

EXPOSE 8080

CMD ["./ecommerce-app"]
