package router

import (
	"aniliststream/internal/middleware"
	"aniliststream/internal/stremio"
	"net/http"
)

func addStremioRoute(mux *http.ServeMux, pattern string, handler http.HandlerFunc) {
	wrapped := middleware.StremioMiddleware(handler)
	mux.HandleFunc("GET "+pattern, wrapped)
	mux.HandleFunc("GET /{config}"+pattern, middleware.ConfigMiddleware(wrapped))
}

func CreateRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	stremioHandler := stremio.NewHandler()

	addStremioRoute(mux, "/manifest.json", stremioHandler.Manifest)
	addStremioRoute(mux, "/logo.png", stremioHandler.Logo)
	addStremioRoute(mux, "/configure", stremioHandler.Index)
	mux.HandleFunc("GET /style.css", stremioHandler.Css)
	mux.HandleFunc("GET /", stremioHandler.Index)

	return mux
}
