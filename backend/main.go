package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my ToDo app")
	})

	log.Fatal(http.ListenAndServe(":3000", nil))
}
