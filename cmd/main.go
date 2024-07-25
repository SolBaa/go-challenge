package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SolBaa/go-challenge/cmd/api"
)

func main() {
	r := api.NewRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "7070"
	}

	log.Printf("Starting server on port %s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), r); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
