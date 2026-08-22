package stremio

import (
	"aniliststream/internal/middleware"
	"aniliststream/internal/models"
	"encoding/json"
	"net/http"
	"strings"
)

func (h *Handler) Catalog(w http.ResponseWriter, r *http.Request) {
	config, exists := middleware.GetConfig(r)
	if !exists {
		http.Error(w, "config not found", http.StatusInternalServerError)
		return
	}
	mediaType := r.PathValue("type")
	id := r.PathValue("id")
	extra := r.PathValue("extra")

	if mediaType != "anime" || !strings.HasPrefix(id, "as:") {
		http.Error(w, "unsupported media type or id", http.StatusBadRequest)
		return
	}

	// In future maybe we can have more catalogs instead of just search
	if extra == "" {
		http.Error(w, "extra is required", http.StatusBadRequest)
		return
	}
	parts := strings.SplitN(extra, "=", 2)

	if len(parts) != 2 || parts[0] != "search" {
		http.Error(w, "unsupported extra", http.StatusBadRequest)
		return
	}
	query := parts[1]
	results, err := h.providerManager.Search(config.Provider, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := models.CatalogResponse{
		Metas: make([]models.MetaPreview, 0, len(results)),
	}

	for _, anime := range results {
		response.Metas = append(response.Metas, models.MetaPreview{
			Id:          "as:" + anime.Id,
			Type:        "anime",
			Name:        anime.Title,
			Poster:      anime.CoverImage,
			Description: anime.Description,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
