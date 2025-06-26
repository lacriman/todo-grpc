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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [title]",
	Short: "delete a new ToDo item",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var id int64
		_, err := fmt.Sscanf(args[0], "%d", &id)
		if err != nil {
			fmt.Printf("invalid id: %v\n", err)
			return
		}

		conn, err := grpc.NewClient("localhost:3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			fmt.Printf("connection error: %v\n", err)
			return
		}
		defer conn.Close()

		client := pb.NewToDoServiceClient(conn)  

		res, err := client.DeleteToDo(context.Background(), &pb.DeleteToDoRequest{
			Id: id,
		})
		if err != nil {
			fmt.Printf("delete failed: %v\n", err)
		}

		fmt.Printf("deleted ToDo: %v\n", res)
	}}
