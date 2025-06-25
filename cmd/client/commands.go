// cmd/client/commands.go
package clientcommands

import (
	"github.com/lacriman/todo-grpc/cmd"
	"github.com/spf13/cobra"
)

var (
	CreateCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a new ToDo item",
		Run: func(cmd *cobra.Command, args []string) {
			// Implementation
		},
	}

	ListCmd = &cobra.Command{
		Use:   "list",
		Short: "List all ToDo items",
		Run: func(cmd *cobra.Command, args []string) {
			// Implementation
		},
	}
)

func init() {
	cmd.RootCmd.AddCommand(CreateCmd)
	cmd.RootCmd.AddCommand(ListCmd)
	// Add other commands here
}
