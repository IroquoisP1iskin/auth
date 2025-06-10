FROM golang:1.23-alpine AS builder

COPY . /app
WORKDIR /app

RUN go mod download
RUN go build -o auth-service cmd/main.go

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/auth-service .

EXPOSE 50051

CMD ["/app/auth-service"]