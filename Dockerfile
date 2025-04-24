# Stage 1: Build Go binary
FROM golang:1.24-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server main.go

# Stage 2: Final image
FROM debian:bookworm-slim

# Install rsvg-convert and dependencies
RUN apt-get update && \
    apt-get install -y librsvg2-bin ca-certificates && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /app/server .
COPY index.html .

EXPOSE 8080

CMD ["./server"]
