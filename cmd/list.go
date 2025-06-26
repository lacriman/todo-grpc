/*
Copyright © 2025 Yaroslav Proskurin <proskurin.yarik@gmail.com>
*/
package cmd

import (
	"context"
	"fmt"

	pb "github.com/lacriman/todo-grpc/pb/todo"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list [title]",
	Short: "list a new ToDo item",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.NewClient("localhost:3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			fmt.Printf("connection error: %v\n", err)
			return
		}
		defer conn.Close()

		client := pb.NewToDoServiceClient(conn)

		res, err := client.ListToDo(context.Background(), &pb.ListToDoRequest{})
		if err != nil {
			fmt.Printf("list failed: %v\n", err)
		}

		fmt.Printf("listed ToDo: %v\n", res)
	}}
