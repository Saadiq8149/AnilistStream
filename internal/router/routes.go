package router

import (
	"aniliststream/internal/middleware"
	"aniliststream/internal/providers"
	"aniliststream/internal/stremio"
	"net/http"
	"strings"
)

func addStremioRoute(mux *http.ServeMux, pattern string, handler http.HandlerFunc) {
	wrapped := middleware.StremioMiddleware(handler)
	mux.HandleFunc("GET "+pattern, wrapped)
	mux.HandleFunc("GET /{config}"+pattern, middleware.ConfigMiddleware(wrapped))
}

func stripJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimSuffix(r.URL.Path, ".json")
		next.ServeHTTP(w, r)
	})
}

func CreateRoutes() http.Handler {
	mux := http.NewServeMux()
	providerManager := providers.NewManager()
	stremioHandler := stremio.NewHandler(providerManager)

	addStremioRoute(mux, "/manifest",
		stremioHandler.Manifest)
	addStremioRoute(mux, "/logo",
		stremioHandler.Logo)
	addStremioRoute(mux, "/configure",
		stremioHandler.Index)

	mux.HandleFunc("GET /{config}/catalog/{type}/{id}/{extra}",
		middleware.ConfigMiddleware(middleware.StremioMiddleware(stremioHandler.Catalog)))
	mux.HandleFunc("GET /{config}/catalog/{type}/{id}",
		middleware.ConfigMiddleware(middleware.StremioMiddleware(stremioHandler.Catalog)))
	mux.HandleFunc("GET /{config}/meta/{type}/{id}",
		middleware.ConfigMiddleware(middleware.StremioMiddleware(stremioHandler.Meta)))
	mux.HandleFunc("GET /{config}/stream/{type}/{id}",
		middleware.ConfigMiddleware(middleware.StremioMiddleware(stremioHandler.Stream)))
	mux.HandleFunc("GET /style.css",
		stremioHandler.Css)
	mux.HandleFunc("GET /",
		stremioHandler.Index)

	return stripJSON(mux)
}
