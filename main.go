package main

import (
	"context"
	"fmt"
	"log"
	"net"

	auth "github.com/frog-in-fog/proto-contracts/gen/auth"

	"google.golang.org/grpc"
)

type authServer struct {
	auth.UnimplementedAuthServiceServer
}

func (s *authServer) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	token := fmt.Sprintf("token-for-%s", req.Username)
	return &auth.LoginResponse{Token: token}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	auth.RegisterAuthServiceServer(grpcServer, &authServer{})

	log.Println("Auth service running on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
