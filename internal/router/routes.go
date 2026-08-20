package router

import (
	"aniliststream/internal/addon"
	"aniliststream/internal/middleware"
	"net/http"
)

func addStremioRoute(mux *http.ServeMux, pattern string, handler http.HandlerFunc) {
	wrapped := middleware.ConfigMiddleware(middleware.StremioMiddleware(handler))
	mux.HandleFunc("GET "+pattern, wrapped)
	mux.HandleFunc("GET /{config}"+pattern, wrapped)
}

func CreateRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	addonHandler := addon.NewHandler()

	addStremioRoute(mux, "/manifest.json", addonHandler.Manifest)
	addStremioRoute(mux, "/logo.png", addonHandler.Logo)
	addStremioRoute(mux, "/", addonHandler.Frontend)

	return mux
}