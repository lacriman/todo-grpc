package cmd

import (
	"log"
	"net"

	"github.com/lacriman/todo-grpc/internal/auth"
	"github.com/lacriman/todo-grpc/internal/handlers"
	pb "github.com/lacriman/todo-grpc/pb/todo"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
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
		reflection.Register(grpcServer)

		log.Println("Server is running on port :3000")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
}
