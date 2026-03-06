package main

import (
	"anilist-stream/internal/anilist"
	"anilist-stream/internal/handlers"
	"anilist-stream/internal/metadata"
	"anilist-stream/internal/streams"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(time.Second * 60))
	r.Use(cors)

	metadataService := metadata.NewMetadataService()
	sourceService := streams.NewSourceService()
	anilistService := anilist.NewAnilistService()
	handlers.RegisterRoutes(r, metadataService, sourceService, anilistService)

	fmt.Println("server running on :" + os.Getenv("PORT"))
	http.ListenAndServe("127.0.0.1:"+os.Getenv("PORT"), r)
}
