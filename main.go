package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/lacriman/todo-grpc/auth"
	"github.com/lacriman/todo-grpc/handlers"
	pb "github.com/lacriman/todo-grpc/pb/todo"
)

func main() {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	jwks, err := auth.LoadJWKS()
	if err != nil {
		log.Fatalf("failed to load JWKS: %v", err)
	}
	defer jwks.EndBackground()

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(auth.AuthInterceptor(jwks)),
	)

	todoServer := handlers.NewServer()
	pb.RegisterToDoServiceServer(grpcServer, todoServer)
	// enables grpcurl
	reflection.Register(grpcServer)

	log.Println("Server is running on port :3000")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
