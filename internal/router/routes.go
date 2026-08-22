package router

import (
	"aniliststream/internal/middleware"
	"aniliststream/internal/providers"
	"aniliststream/internal/stremio"
	"net/http"
	"strings"
)

func addStremioRoute(url string, handler http.HandlerFunc, mux *http.ServeMux) {
	wrapped := middleware.StremioMiddleware(handler)
	mux.HandleFunc(url, middleware.ConfigMiddleware(wrapped))
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

	mux.HandleFunc("GET /manifest", middleware.StremioMiddleware(stremioHandler.Manifest))
	addStremioRoute("GET /{config}/manifest", stremioHandler.Manifest, mux)
	mux.HandleFunc("GET /logo", middleware.StremioMiddleware(stremioHandler.Logo))
	addStremioRoute("GET /{config}/logo", stremioHandler.Logo, mux)
	mux.HandleFunc("GET /configure", middleware.StremioMiddleware(stremioHandler.Index))
	addStremioRoute("GET /{config}/configure", stremioHandler.Index, mux)

	addStremioRoute("GET /{config}/catalog/{type}/{id}/{extra}", stremioHandler.Catalog, mux)
	addStremioRoute("GET /{config}/catalog/{type}/{id}", stremioHandler.Catalog, mux)
	addStremioRoute("GET /{config}/meta/{type}/{id}", stremioHandler.Meta, mux)
	addStremioRoute("GET /{config}/stream/{type}/{id}", stremioHandler.Stream, mux)

	mux.HandleFunc("GET /style.css", stremioHandler.Css)
	mux.HandleFunc("GET /", stremioHandler.Index)

	return stripJSON(mux)
}
