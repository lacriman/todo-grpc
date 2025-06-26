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

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [title]",
	Short: "Create a new ToDo item",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]

		conn, err := grpc.NewClient("localhost:3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			fmt.Printf("connection error: %v\n", err)
			return
		}
		defer conn.Close()

		client := pb.NewToDoServiceClient(conn)

		res, err := client.CreateToDo(context.Background(), &pb.CreateToDoRequest{
			Title: title,
		})
		if err != nil {
			fmt.Printf("create failed: %v\n", err)
		}

		fmt.Printf("Created ToDo: %v\n", res.Todo.Title)
	}}
