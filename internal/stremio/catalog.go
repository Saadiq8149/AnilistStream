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

func (s *StremioHandler) CatalogHandler(w http.ResponseWriter, r *http.Request) {
	extra := chi.URLParam(r, "extra")

	if extra != "" {
		extra = strings.TrimSuffix(extra, ".json")
		searchQuery := strings.TrimPrefix(extra, "search=")

		anime, err := s.MetadataService.Provider.SearchAnime(searchQuery)
		if err != nil {
			http.Error(w, "Error searching for anime", http.StatusInternalServerError)
			return
		}

		if len(anime) > 25 {
			anime = anime[:25]
		}

		var metas []types.MetaPreview

		for _, a := range anime {
			release := ""
			if a.FromYear > 0 && a.ToYear > 0 {
				release = fmt.Sprintf("%d-%d", a.FromYear, a.ToYear)
			} else if a.FromYear > 0 {
				release = fmt.Sprintf("%d-", a.FromYear)
			}

			description := util.StripHTML(a.Description)
			meta := types.MetaPreview{
				ID:          fmt.Sprintf("ani_%s_%s_%s", a.AnilistID, a.ProviderID, a.MalID),
				Type:        "series",
				Name:        a.Title,
				Poster:      a.Poster,
				Genres:      a.Genres,
				IMDBRating:  fmt.Sprintf("%.1f", a.Rating),
				ReleaseInfo: release,
				Description: description,
			}

			metas = append(metas, meta)
		}

		response := types.CatalogResponse{
			Metas: metas,
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Cache-Control", "max-age=3600")

		json.NewEncoder(w).Encode(response)
		return
	}

	json.NewEncoder(w).Encode(types.CatalogResponse{
		Metas: []types.MetaPreview{},
	})
}
