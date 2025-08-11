package main

import (
	"context"
	"log"
	"time"

	auth "github.com/frog-in-fog/proto-contracts/gen/auth"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := auth.NewAuthServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	resp, err := client.Login(ctx, &auth.LoginRequest{
		Username: "testuser",
		Password: "1234",
	})
	if err != nil {
		log.Fatalf("failed to login: %v", err)
	}

	log.Printf("Received token: %s", resp.Token)
}
