package anilist

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const anilistAPI = "https://graphql.anilist.co"

type AnilistService struct{}

func NewAnilistService() *AnilistService {
	return &AnilistService{}
}

func (s *AnilistService) SyncProgress(anilistID string, episode int, accessToken string) error {
	query := `
	query ($mediaId: Int) {
		Media(id: $mediaId) {
			episodes
			status
			nextAiringEpisode {
                episode
            }
			mediaListEntry {
				progress
				status
				repeat
			}
		}
	}`

	variables := map[string]any{
		"mediaId": anilistID,
	}

	body := map[string]any{
		"query":     query,
		"variables": variables,
	}

	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", anilistAPI, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var result struct {
		Data struct {
			Media struct {
				Episodes          int    `json:"episodes"`
				Status            string `json:"status"`
				NextAiringEpisode struct {
					Episode int `json:"episode"`
				} `json:"nextAiringEpisode"`
				MediaListEntry struct {
					Progress int    `json:"progress"`
					Status   string `json:"status"`
					Repeat   int    `json:"repeat"`
				} `json:"mediaListEntry"`
			} `json:"Media"`
		} `json:"data"`
	}

	json.NewDecoder(resp.Body).Decode(&result)

	state := result.Data.Media

	status := state.MediaListEntry.Status

	totalEpisodes := state.Episodes
	if state.Status != "FINISHED" {
		if state.NextAiringEpisode.Episode > 0 {
			totalEpisodes = state.NextAiringEpisode.Episode - 1
		} else {
			totalEpisodes = 0
		}
	}

	switch status {
	case "COMPLETED":
		if episode < totalEpisodes {
			status = "REPEATING"
		}
	case "PLANNING", "PAUSED", "DROPPED":
		status = "CURRENT"
	}

	if totalEpisodes > 0 && episode >= totalEpisodes && state.Status == "FINISHED" {
		status = "COMPLETED"
	}

	mutation := `
	mutation ($mediaId: Int, $progress: Int, $status: MediaListStatus) {
		SaveMediaListEntry(
			mediaId: $mediaId
			progress: $progress
			status: $status
		) {
			id
		}
	}`

	vars := map[string]any{
		"mediaId":  anilistID,
		"progress": episode,
		"status":   status,
	}

	body = map[string]any{
		"query":     mutation,
		"variables": vars,
	}

	jsonBody, _ = json.Marshal(body)

	req, _ = http.NewRequest("POST", anilistAPI, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
