package handlers

import (
	"anilist-stream/internal/anilist"
	"anilist-stream/internal/metadata"
	"anilist-stream/internal/pages"
	"anilist-stream/internal/streams"
	"anilist-stream/internal/stremio"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router, metadataService *metadata.MetadataService, sourceService *streams.SourceService, anilistService *anilist.AnilistService) {
	s := stremio.NewStremioHandler(metadataService, sourceService, anilistService)

	r.Get("/catalog/{type}/{id}.json", s.CatalogHandler)
	r.Get("/{anilist_token}/catalog/{type}/{id}.json", s.CatalogHandler)
	r.Get("/catalog/{type}/{id}/{extra}.json", s.CatalogHandler)
	r.Get("/{anilist_token}/catalog/{type}/{id}/{extra}.json", s.CatalogHandler)

	r.Get("/meta/{type}/{id}.json", s.MetaHandler)
	r.Get("/{anilist_token}/meta/{type}/{id}.json", s.MetaHandler)

	r.Get("/stream/{type}/{id}.json", s.StreamHandler)
	r.Get("/{anilist_token}/stream/{type}/{id}.json", s.StreamHandler)

	r.Handle("/logo.png", http.FileServer(http.Dir("./public")))
	r.Handle("/manifest.json", http.FileServer(http.Dir("./public")))

	r.Get("/", pages.IndexHandler)
	r.Get("/{anilist_token}", pages.IndexHandler)
	r.Get("/configure", pages.ConfigureHandler)
	r.Get("/{anilist_token}/configure", pages.ConfigureHandler)
	r.Get("/{anilist_token}/manifest.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/manifest.json")
	})
}
