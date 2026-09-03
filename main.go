package main

import (
	"aniliststream/internal/router"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	fmt.Printf("server running on address: %s\n", os.Getenv("BACKEND_ADDR"))

	server := router.NewServer()
	if err := server.Start(); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
