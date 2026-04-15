package metadata

import (
	"anilist-stream/internal/types"
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type AllAnimeProvider struct{}

func NewAllAnimeProvider() *AllAnimeProvider {
	return &AllAnimeProvider{}
}

func (a *AllAnimeProvider) Name() string {
	return "AllAnime"
}

const (
	allanimeAPI     = "https://api.allanime.day"
	allanimeReferer = "https://allmanga.to"
	userAgent       = "Mozilla/5.0"
)

func (p *AllAnimeProvider) SearchAnime(query string) ([]types.Metadata, error) {
	searchGql := `
	query (
		$search: SearchInput
		$limit: Int
		$translationType: VaildTranslationTypeEnumType
		$countryOrigin: VaildCountryOriginEnumType
	) {
		shows(
			search: $search
			limit: $limit
			page: 1
			translationType: $translationType
			countryOrigin: $countryOrigin
		) {
			edges {
				_id
				name
				englishName
				nativeName
				description
				thumbnail
				banner
				genres
				score
				aniListId
				malId
				episodeCount
				airedStart
				airedEnd
				status
			}
		}
	}`

	requestBody := map[string]any{
		"query": searchGql,
		"variables": map[string]any{
			"search": map[string]any{
				"allowAdult":   true,
				"allowUnknown": false,
				"query":        strings.ToLower(query),
			},
			"limit":           40,
			"translationType": "sub",
			"countryOrigin":   "ALL",
		},
	}

	bodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", allanimeAPI+"/api", bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Referer", allanimeReferer)
	req.Header.Set("User-Agent", userAgent)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Data struct {
			Shows struct {
				Edges []struct {
					ID          string   `json:"_id"`
					Name        string   `json:"name"`
					EnglishName string   `json:"englishName"`
					NativeName  string   `json:"nativeName"`
					Description string   `json:"description"`
					Thumbnail   string   `json:"thumbnail"`
					Banner      string   `json:"banner"`
					Genres      []string `json:"genres"`
					Score       float64  `json:"score"`
					AniListID   string   `json:"aniListId"`
					MalID       string   `json:"malId"`
					EpisodeCnt  string   `json:"episodeCount"`
					Status      string   `json:"status"`
				} `json:"edges"`
			} `json:"shows"`
		} `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	var metadata []types.Metadata

	for _, anime := range result.Data.Shows.Edges {
		title := anime.EnglishName
		if title == "" {
			title = anime.Name
		}

		episodes := 0
		if anime.EpisodeCnt != "" {
			episodes, _ = strconv.Atoi(anime.EpisodeCnt)
		}

		metadata = append(metadata, types.Metadata{
			ProviderID:  anime.ID,
			AnilistID:   anime.AniListID,
			MalID:       anime.MalID,
			Title:       title,
			Description: anime.Description,
			Poster:      anime.Thumbnail,
			Banner:      anime.Banner,
			Rating:      anime.Score,
			Episodes:    episodes,
			Genres:      anime.Genres,
			Status:      anime.Status,
		})
	}

	return metadata, nil
}

func (p *AllAnimeProvider) GetAnime(id string) (types.Metadata, error) {
	query := `
	query ($id: String!) {
		show(_id: $id) {
			_id
			name
			englishName
			nativeName
			description
			thumbnail
			banner
			genres
			score
			aniListId
			malId
			episodeCount
			airedStart
			airedEnd
			status
		}
	}`

	requestBody := map[string]any{
		"query": query,
		"variables": map[string]any{
			"id": id,
		},
	}

	bodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return types.Metadata{}, err
	}

	req, err := http.NewRequest("POST", allanimeAPI+"/api", bytes.NewBuffer(bodyBytes))
	if err != nil {
		return types.Metadata{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Referer", allanimeReferer)
	req.Header.Set("User-Agent", userAgent)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return types.Metadata{}, err
	}
	defer resp.Body.Close()

	var result struct {
		Data struct {
			Show struct {
				ID          string   `json:"_id"`
				Name        string   `json:"name"`
				EnglishName string   `json:"englishName"`
				NativeName  string   `json:"nativeName"`
				Description string   `json:"description"`
				Thumbnail   string   `json:"thumbnail"`
				Banner      string   `json:"banner"`
				Genres      []string `json:"genres"`
				Score       float64  `json:"score"`
				AniListID   string   `json:"aniListId"`
				MalID       string   `json:"malId"`
				EpisodeCnt  string   `json:"episodeCount"`
				Status      string   `json:"status"`
			} `json:"show"`
		} `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return types.Metadata{}, err
	}

	anime := result.Data.Show

	title := anime.EnglishName
	if title == "" {
		title = anime.Name
	}

	episodes := 0
	if anime.EpisodeCnt != "" {
		episodes, _ = strconv.Atoi(anime.EpisodeCnt)
	}

	return types.Metadata{
		ProviderID:  anime.ID,
		AnilistID:   anime.AniListID,
		MalID:       anime.MalID,
		Title:       title,
		Description: anime.Description,
		Poster:      anime.Thumbnail,
		Banner:      anime.Banner,
		Rating:      anime.Score,
		Episodes:    episodes,
		Genres:      anime.Genres,
		Status:      anime.Status,
	}, nil
}
