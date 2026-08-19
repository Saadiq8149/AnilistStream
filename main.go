package main

import (
	"aniliststream/internal/router"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	fmt.Printf("server running on address: %s\n", os.Getenv("BACKEND_ADDR"))

	server := router.NewServer()
	if err := server.Start(); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
