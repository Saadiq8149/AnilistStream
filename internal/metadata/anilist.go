package metadata

import (
	"anilist-stream/internal/types"
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type AnilistProvider struct {
	client *http.Client
}

func NewAnilistProvider() *AnilistProvider {
	return &AnilistProvider{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (a *AnilistProvider) Name() string {
	return "Anilist"
}

func (a *AnilistProvider) SearchAnime(query string) ([]types.Metadata, error) {
	graphqlQuery := `
	query ($search: String) {
		Page(page: 1, perPage: 50) {
			media(search: $search, type: ANIME) {
				id
				idMal
				status
				episodes
				nextAiringEpisode {
					episode
				}
				title {
					romaji
					english
				}
				description(asHtml: false)
				coverImage {
					large
				}
				bannerImage
				startDate {
					year
				}
				endDate {
					year
				}
				averageScore
				genres
			}
		}
	}`

	requestBody := map[string]any{
		"query": graphqlQuery,
		"variables": map[string]any{
			"search": query,
		},
	}

	bodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://graphql.anilist.co", bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Data struct {
			Page struct {
				Media []struct {
					ID     json.Number `json:"id"`
					IDMal  json.Number `json:"idMal"`
					Status string      `json:"status"`

					Episodes int `json:"episodes"`

					NextAiringEpisode *struct {
						Episode int `json:"episode"`
					} `json:"nextAiringEpisode"`

					Title struct {
						Romaji  string `json:"romaji"`
						English string `json:"english"`
					} `json:"title"`

					Description string `json:"description"`

					CoverImage struct {
						Large string `json:"large"`
					} `json:"coverImage"`

					BannerImage string `json:"bannerImage"`

					StartDate struct {
						Year int `json:"year"`
					} `json:"startDate"`

					EndDate struct {
						Year int `json:"year"`
					} `json:"endDate"`

					AverageScore int      `json:"averageScore"`
					Genres       []string `json:"genres"`
				} `json:"media"`
			} `json:"Page"`
		} `json:"data"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	var metadata []types.Metadata

	for _, anime := range result.Data.Page.Media {

		title := anime.Title.English
		if title == "" {
			title = anime.Title.Romaji
		}

		episodes := anime.Episodes

		if anime.Status != "FINISHED" {
			if anime.NextAiringEpisode != nil && anime.NextAiringEpisode.Episode > 0 {
				episodes = anime.NextAiringEpisode.Episode - 1
			} else {
				episodes = 0
			}
		}

		metadata = append(metadata, types.Metadata{
			ProviderID:  anime.ID.String(),
			AnilistID:   anime.ID.String(),
			MalID:       anime.IDMal.String(),
			Title:       title,
			Description: anime.Description,
			Poster:      anime.CoverImage.Large,
			Banner:      anime.BannerImage,
			FromYear:    anime.StartDate.Year,
			ToYear:      anime.EndDate.Year,
			Rating:      float64(anime.AverageScore) / 10,
			Episodes:    episodes,
			Genres:      anime.Genres,
			Status:      anime.Status,
		})
	}

	return metadata, nil
}

func (a *AnilistProvider) GetAnime(id string) (types.Metadata, error) {
	graphqlQuery := `
	query ($id: Int) {
		Media(id: $id, type: ANIME) {
			id
			idMal
			status
			episodes
			nextAiringEpisode {
				episode
			}
			title {
				romaji
				english
			}
			description(asHtml: false)
			coverImage {
				large
			}
			bannerImage
			startDate {
				year
			}
			endDate {
				year
			}
			averageScore
			genres
		}
	}`

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return types.Metadata{}, err
	}

	requestBody := map[string]any{
		"query": graphqlQuery,
		"variables": map[string]any{
			"id": idInt,
		},
	}

	bodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return types.Metadata{}, err
	}

	req, err := http.NewRequest("POST", "https://graphql.anilist.co", bytes.NewBuffer(bodyBytes))
	if err != nil {
		return types.Metadata{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := a.client.Do(req)
	if err != nil {
		return types.Metadata{}, err
	}
	defer resp.Body.Close()

	var result struct {
		Data struct {
			Media struct {
				ID     json.Number `json:"id"`
				IDMal  json.Number `json:"idMal"`
				Status string      `json:"status"`

				Episodes int `json:"episodes"`

				NextAiringEpisode *struct {
					Episode int `json:"episode"`
				} `json:"nextAiringEpisode"`

				Title struct {
					Romaji  string `json:"romaji"`
					English string `json:"english"`
				} `json:"title"`

				Description string `json:"description"`

				CoverImage struct {
					Large string `json:"large"`
				} `json:"coverImage"`

				BannerImage string `json:"bannerImage"`

				StartDate struct {
					Year int `json:"year"`
				} `json:"startDate"`

				EndDate struct {
					Year int `json:"year"`
				} `json:"endDate"`

				AverageScore int      `json:"averageScore"`
				Genres       []string `json:"genres"`
			} `json:"Media"`
		} `json:"data"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return types.Metadata{}, err
	}

	anime := result.Data.Media

	title := anime.Title.English
	if title == "" {
		title = anime.Title.Romaji
	}

	episodes := anime.Episodes

	if anime.Status != "FINISHED" {
		if anime.NextAiringEpisode != nil && anime.NextAiringEpisode.Episode > 0 {
			episodes = anime.NextAiringEpisode.Episode - 1
		} else {
			episodes = 0
		}
	}

	meta := types.Metadata{
		ProviderID:  anime.ID.String(),
		AnilistID:   anime.ID.String(),
		MalID:       anime.IDMal.String(),
		Title:       title,
		Description: anime.Description,
		Poster:      anime.CoverImage.Large,
		Banner:      anime.BannerImage,
		FromYear:    anime.StartDate.Year,
		ToYear:      anime.EndDate.Year,
		Rating:      float64(anime.AverageScore) / 10,
		Episodes:    episodes,
		Genres:      anime.Genres,
		Status:      anime.Status,
	}

	return meta, nil
}
