package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/lacriman/todo-grpc/handlers"
	pb "github.com/lacriman/todo-grpc/pb/todo"
)

func main() {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	todoServer := handlers.NewServer()
	pb.RegisterToDoServiceServer(grpcServer, todoServer)

	log.Println("Server is running on port :3000")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
