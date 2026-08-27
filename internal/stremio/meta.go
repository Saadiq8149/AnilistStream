package stremio

import (
	"aniliststream/internal/middleware"
	"aniliststream/internal/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
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

	anilistId, anilistIdExists := anime.ExternalIds["anilist"]

	response := models.MetaResponse{
		Meta: models.Meta{
			Id:          id,
			Type:        "anime",
			Name:        anime.Title,
			Genres:      anime.Genres,
			Poster:      anime.CoverImage,
			Background:  anime.BannerImage,
			Description: anime.Description,
			ImdbRating:  anime.Rating,
		},
	}

	if anime.Year != 0 {
		response.Meta.ReleaseInfo = fmt.Sprintf("%d", anime.Year)
	}

	var kitsuResponse struct {
		Meta models.Meta `json:"meta"`
	}
	kitsuId, kitsuExists := anime.ExternalIds["kitsu"]
	if kitsuExists {
		kitsuExists = false
		kitsuURL := "https://anime-kitsu.strem.fun/meta/series/kitsu%3A" + kitsuId + ".json"
		req, err := http.NewRequest(
			http.MethodGet,
			kitsuURL,
			nil,
		)
		if err == nil {
			req.Header.Set("Referer", "https://web.stremio.com/")
			req.Header.Set("User-Agent",
				"Mozilla/5.0 (X11; Linux x86_64; rv:154.0) Gecko/20100101 Firefox/154.0",
			)

			resp, err := http.DefaultClient.Do(req)
			if err == nil {
				defer resp.Body.Close()

				body, err := io.ReadAll(resp.Body)
				if err == nil && resp.StatusCode == http.StatusOK {
					if err := json.Unmarshal(body, &kitsuResponse); err == nil {
						response.Meta.Logo = kitsuResponse.Meta.Logo
						response.Meta.Background = kitsuResponse.Meta.Background
						response.Meta.Runtime = kitsuResponse.Meta.Runtime
						kitsuExists = true
					}
				}
			} else {
				fmt.Println(err)
			}
		}
	}

	minEpNumber := 0
	for _, ep := range anime.Episodes {
		if minEpNumber == 0 || ep.Number < minEpNumber {
			minEpNumber = ep.Number
		}
	}
	for _, episode := range anime.Episodes {
		id := "as:" + animeID + ":" + episode.Id + ":" + strconv.Itoa(episode.Number)
		if anilistIdExists {
			id = "as:" + animeID + ":" + episode.Id + ":" + strconv.Itoa(episode.Number) + ":" + anilistId
		}

		title := episode.Title
		thumbnail := episode.Thumbnail
		released := episode.Released
		episodeNumber := episode.Number - minEpNumber + 1

		if kitsuExists {
			for _, kitsuEpisode := range kitsuResponse.Meta.Videos {
				if kitsuEpisode.Episode == episodeNumber {
					title = kitsuEpisode.Title
					thumbnail = kitsuEpisode.Thumbnail
					released = kitsuEpisode.Released
					break
				}
			}
		}

		response.Meta.Videos = append(response.Meta.Videos, models.Video{
			Id:        id,
			Title:     title,
			Episode:   episodeNumber,
			Thumbnail: thumbnail,
			Released:  released,
			Season:    1,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
