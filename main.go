package main

import (
	"anilist-stream/internal/handlers"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(time.Second * 60))
	handlers.RegisterRoutes(r)

	fmt.Println("server running on Port: " + os.Getenv("PORT"))
	http.ListenAndServe("127.0.0.1:"+os.Getenv("PORT"), r)
}
