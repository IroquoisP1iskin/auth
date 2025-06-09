package main

import (
	"context"
	"fmt"
	"github.com/IroquoisP1iskin/auth/pkg/note_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
)

const (
	grpcPort = 50051
)

type server struct {
	note_v1.UnimplementedAuthServer
}

// GetNote implements note_v1.AuthServer
func (s *server) Get(ctx context.Context, req *note_v1.GetRequest) (*note_v1.GetResponse, error) {
	log.Printf("GetNote called with ID: %d", req.Id)
	return &note_v1.GetResponse{
		Id:        23,
		Name:      "Solid Snake",
		Email:     "bitch@suckmycock.dildo",
		Role:      note_v1.Role_ADMIN,
		CreatedAt: timestamppb.Now(),
		UpdatedAt: timestamppb.Now(),
	}, nil
}
func main() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort)) // прочитать про порт 127.00.1:50051
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	serv := grpc.NewServer()
	reflection.Register(serv) // регистрируем рефлексию для gRPC сервера - зачем это нужно?

	note_v1.RegisterAuthServer(serv, &server{})
	log.Printf("server listening at %v", listen.Addr())
	if err := serv.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
