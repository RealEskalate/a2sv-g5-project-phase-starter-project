FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o bin/blog-api ./cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/bin/blog-api .
EXPOSE 8081
CMD ["./blog-api"]
