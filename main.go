package main

import (
	"log"

	"github.com/lacriman/todo-grpc/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
