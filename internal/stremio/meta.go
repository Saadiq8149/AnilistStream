package stremio

import (
	"aniliststream/internal/middleware"
	"aniliststream/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (h *Handler) Meta(w http.ResponseWriter, r *http.Request) {
	config, exists := middleware.GetConfig(r)
	if !exists {
		http.Error(w, "config not found", http.StatusInternalServerError)
		return
	}

	mediaType := r.PathValue("type")
	id := r.PathValue("id")

	if mediaType != "anime" || !strings.HasPrefix(id, "as:") {
		http.Error(w, "unsupported media type or id", http.StatusBadRequest)
		return
	}

	animeID := strings.TrimPrefix(id, "as:")

	anime, err := h.providerManager.GetDetails(config.Provider, animeID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := models.MetaResponse{
		Meta: models.Meta{
			Id:          id,
			Type:        "anime",
			Name:        anime.Title,
			Genres:      anime.Genres,
			Poster:      anime.CoverImage,
			Background:  anime.BannerImage,
			Description: anime.Description,
		},
	}

	if anime.Year != 0 {
		response.Meta.ReleaseInfo = fmt.Sprintf("%d", anime.Year)
	}

	for _, episode := range anime.Episodes {
		response.Meta.Videos = append(response.Meta.Videos, models.Video{
			Id:        "as:" + animeID + ":" + episode.Id,
			Title:     episode.Title,
			Episode:   episode.Number,
			Thumbnail: episode.Thumbnail,
			Released:  episode.Released,
			Season:    1,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
