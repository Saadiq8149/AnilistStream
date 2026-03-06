package stremio

import (
	"anilist-stream/internal/types"
	"anilist-stream/internal/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

func (s *StremioHandler) MetaHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	idParam = strings.TrimSuffix(idParam, ".json")

	parts := strings.Split(idParam, "_")
	if len(parts) != 3 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	id := parts[2]

	anime, err := s.MetadataService.Provider.GetAnime(id)
	if err != nil {
		http.Error(w, "Error fetching anime", http.StatusInternalServerError)
		return
	}

	release := ""
	if anime.FromYear > 0 && anime.ToYear > 0 {
		release = fmt.Sprintf("%d-%d", anime.FromYear, anime.ToYear)
	} else if anime.FromYear > 0 {
		release = fmt.Sprintf("%d-", anime.FromYear)
	}

	description := util.StripHTML(anime.Description)

	var videos []types.Video

	for i := 1; i <= anime.Episodes; i++ {

		video := types.Video{
			ID:       fmt.Sprintf("%s:%d", idParam, i),
			Title:    fmt.Sprintf("Episode %d", i),
			Episode:  i,
			Season:   1,
			Released: fmt.Sprintf("%d-01-01T00:00:00.000Z", anime.FromYear),
		}

		videos = append(videos, video)
	}

	meta := types.Meta{
		ID:          idParam,
		Type:        "series",
		Name:        anime.Title,
		Genres:      anime.Genres,
		Poster:      anime.Poster,
		Background:  anime.Banner,
		Description: description,
		ReleaseInfo: release,
		IMDBRating:  fmt.Sprintf("%.1f", anime.Rating),
		Videos:      videos,
	}

	response := types.MetaResponse{
		Meta: meta,
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "max-age=3600")

	json.NewEncoder(w).Encode(response)
}
