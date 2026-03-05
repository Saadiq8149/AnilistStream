package handlers

import (
	"anilist-stream/internal/stremio"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router) {
	r.Get("/manifest.json", stremio.ManifestHandler)
	r.Get("/catalog/{type}/{id}.json", stremio.CatalogHandler)
	r.Get("/catalog/{type}/{id}/{extra}.json", stremio.CatalogHandler)
}
