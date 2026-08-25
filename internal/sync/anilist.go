package sync

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const anilistAPI = "https://graphql.anilist.co"

func SyncAnilist(token string, anilistId string, episodeNumber int) {
	mediaID, err := strconv.Atoi(anilistId)
	if err != nil {
		fmt.Println("invalid anilist id:", err)
		return
	}

	query := `
		query ($id: Int!) {
			Media(id: $id, type: ANIME) {
				status
				episodes
				mediaListEntry {
					id
					status
					progress
				}
			}
		}
	`

	variables := map[string]any{
		"id": mediaID,
	}

	payload := map[string]any{
		"query":     query,
		"variables": variables,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("failed to marshal anilist query:", err)
		return
	}

	req, err := http.NewRequest(
		http.MethodPost,
		anilistAPI,
		bytes.NewReader(body),
	)
	if err != nil {
		fmt.Println("failed to create anilist request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("anilist request failed:", err)
		return
	}
	defer resp.Body.Close()

	var result struct {
		Data struct {
			Media struct {
				Status   string `json:"status"`
				Episodes int    `json:"episodes"`
				List     *struct {
					ID       int    `json:"id"`
					Status   string `json:"status"`
					Progress int    `json:"progress"`
				} `json:"mediaListEntry"`
			} `json:"Media"`
		} `json:"data"`
		Errors []struct {
			Message string `json:"message"`
		} `json:"errors"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println("failed to decode AniList response:", err)
		return
	}

	if len(result.Errors) > 0 {
		fmt.Println("AniList error:", result.Errors[0].Message)
		return
	}

	media := result.Data.Media
	list := media.List

	status := ""

	if list != nil {
		switch list.Status {
		case "PLANNING":
			status = "CURRENT"
		case "CURRENT":
			status = "CURRENT"
		case "REPEATING":
			status = "REPEATING"
		case "PAUSED":
			status = "CURRENT"
		case "DROPPED":
			status = "CURRENT"
		case "COMPLETED":
			status = "REPEATING"
		}
	} else {
		status = "CURRENT"
	}

	if media.Status == "FINISHED" &&
		media.Episodes > 0 &&
		episodeNumber >= media.Episodes {
		status = "COMPLETED"
	}

	mutation := `
		mutation (
			$mediaId: Int!,
			$progress: Int!,
			$status: MediaListStatus
		) {
			SaveMediaListEntry(
				mediaId: $mediaId
				progress: $progress
				status: $status
			) {
				id
				status
				progress
			}
		}
	`

	mutationPayload := map[string]any{
		"query": mutation,
		"variables": map[string]any{
			"mediaId":  mediaID,
			"progress": episodeNumber,
			"status":   status,
		},
	}

	body, err = json.Marshal(mutationPayload)
	if err != nil {
		fmt.Println("failed to marshal AniList mutation:", err)
		return
	}

	req, err = http.NewRequest(
		http.MethodPost,
		anilistAPI,
		bytes.NewReader(body),
	)
	if err != nil {
		fmt.Println("failed to create AniList mutation request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("AniList mutation failed:", err)
		return
	}
	defer resp.Body.Close()
}
