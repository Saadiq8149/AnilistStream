package stremio

import (
	"anilist-stream/internal/types"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

func (s *StremioHandler) StreamHandler(w http.ResponseWriter, r *http.Request) {
	anilistToken := chi.URLParam(r, "anilist_token")

	idParam := chi.URLParam(r, "id")
	idParam = strings.TrimSuffix(idParam, ".json")

	parts := strings.Split(idParam, "%3A")

	if len(parts) != 2 {
		http.Error(w, "Invalid stream ID", http.StatusBadRequest)
		return
	}

	animeID := parts[0]
	episodeStr := parts[1]

	parts = strings.Split(strings.TrimPrefix(animeID, "ani_"), "_")
	anilistID := parts[0]
	// providerID := parts[1]
	malID := parts[2]

	episode, err := strconv.Atoi(episodeStr)
	if err != nil {
		http.Error(w, "Invalid episode", http.StatusBadRequest)
		return
	}

	if anilistToken != "" {
		s.AnilistService.SyncProgress(anilistID, episode, anilistToken)
	}

	sources, err := s.SourceService.GetStreams(anilistID, malID, episode)
	if err != nil {
		http.Error(w, "Source fetch failed", http.StatusInternalServerError)
		return
	}

	var streams []types.Stream

	for _, src := range sources {
		stream := types.Stream{
			Name:      src.Name,
			Title:     src.Name,
			Url:       src.Url,
			Subtitles: src.Subtitles,
		}

		if src.IsHLS {
			stream.BehaviorHints = &types.BehaviorHints{
				NotWebReady: true,
			}
		}

		if len(src.Headers) > 0 {
			stream.BehaviorHints = &types.BehaviorHints{
				NotWebReady: true,
				ProxyHeaders: map[string]map[string]string{
					"request": src.Headers,
				},
			}
		}

		streams = append(streams, stream)
	}

	response := types.StreamResponse{
		Streams: streams,
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "max-age=60")

	json.NewEncoder(w).Encode(response)
}
