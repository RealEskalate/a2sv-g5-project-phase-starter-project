FROM golang:1.23.0-bullseye AS builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./
RUN go build -o main ./delivery

CMD ["./main"]

FROM debian:bookworm-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/main /app/main
COPY --from=builder /app/.env /app/.env

CMD ["./main"]
