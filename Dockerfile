FROM golang:1.23-alpine AS builder

COPY . /github.com/IroquoisP1iskin/auth/source
WORKDIR /github.com/IroquoisP1iskin/auth/source

RUN go mod download
RUN go build -o /source/bin/auth-service cmd/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/IroquoisP1iskin/auth/source/bin/auth-service .

CMD ["./auth-service"]