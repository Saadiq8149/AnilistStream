package stremio

import (
	"aniliststream/internal/middleware"
	"aniliststream/internal/models"
	"encoding/json"
	"net/http"
	"strings"
)

func (h *Handler) Stream(w http.ResponseWriter, r *http.Request) {
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

	parts := strings.Split(id, ":")
	if len(parts) != 3 {
		http.Error(w, "invalid stream id", http.StatusBadRequest)
		return
	}

	animeID := parts[1]
	episodeID := parts[2]

	sources, err := h.providerManager.GetStreams(
		config.Provider,
		animeID,
		episodeID,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := models.StreamResponse{
		Streams: make([]models.Stream, 0, len(sources)),
	}

	for _, source := range sources {
		response.Streams = append(response.Streams, models.Stream{
			Name:        source.Name,
			Url:         source.Url,
			Description: source.SourceType,
			Subtitles:   source.Subtitles,
			BehaviorHints: &models.BehaviorHints{
				BingeGroup: "aniliststream-" + source.Name,
			},
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
