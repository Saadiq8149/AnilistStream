package stremio

import (
	"anilist-stream/internal/types"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
)

func (s *StremioHandler) StreamHandler(w http.ResponseWriter, r *http.Request) {
	anilistToken := chi.URLParam(r, "anilist_token")

	idParam := chi.URLParam(r, "id")
	idParam = strings.TrimSuffix(idParam, ".json")

	var anilistID string
	var episode int
	var malID string
	var err error
	// var providerID string

	if strings.HasPrefix(idParam, "kitsu") {
		parts := strings.Split(strings.TrimPrefix(idParam, "kitsu%3A"), "%3A")
		if len(parts) != 2 {
			http.Error(w, "Invalid stream ID", http.StatusBadRequest)
			return
		}

		kitsuID := parts[0]
		episodeStr := parts[1]

		episode, err = strconv.Atoi(episodeStr)
		if err != nil {
			http.Error(w, "Invalid episode", http.StatusBadRequest)
			return
		}

		// first check Redis cache for AniList ID mapping
		cachedAniListID, err := s.RedisService.Get("idmap:kitsu:" + kitsuID)
		if err == nil && cachedAniListID != "" {
			anilistID = cachedAniListID
		} else {
			idMap, err := s.IDMapService.GetIDMap(kitsuID, "kitsu")
			if err != nil {
				http.Error(w, "ID mapping failed", http.StatusInternalServerError)
				return
			}

			anilistID = idMap["anilist"]
			if anilistID == "" {
				http.Error(w, "ID mapping not found", http.StatusNotFound)
				return
			}
			s.RedisService.Set("idmap:kitsu:"+kitsuID, anilistID, 90*24*time.Hour)
		}

	} else {
		parts := strings.Split(idParam, "%3A")

		if len(parts) != 2 {
			http.Error(w, "Invalid stream ID", http.StatusBadRequest)
			return
		}

		animeID := parts[0]
		episodeStr := parts[1]

		parts = strings.Split(strings.TrimPrefix(animeID, "ani_"), "_")
		anilistID = parts[0]
		// providerID = parts[1]
		malID = parts[2]

		episode, err = strconv.Atoi(episodeStr)
		if err != nil {
			http.Error(w, "Invalid episode", http.StatusBadRequest)
			return
		}
	}

	if anilistToken != "" {
		// to prevent syncing the prefetched episode by stemio, which user hasn't actually watched
		skipSync, err := s.RedisService.Exists("sync:" + anilistID)
		if err != nil {
			http.Error(w, "Redis error", http.StatusInternalServerError)
			return
		}

		if !skipSync {
			s.AnilistService.SyncProgress(anilistID, episode, anilistToken)
			s.RedisService.Set("sync:"+anilistID, episode, 1*time.Second)
		}
	}
	cacheKey := "streams:" + anilistID + ":" + malID + ":" + strconv.Itoa(episode)

	var cached types.StreamResponse
	found, err := s.RedisService.GetJSON(cacheKey, &cached)
	if err == nil && found {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cached)
		return
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

	if len(streams) > 0 {
		ttl := 30*time.Minute + time.Duration(rand.Intn(120))*time.Second
		s.RedisService.SetJSON(cacheKey, response, ttl)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "max-age=60")

	json.NewEncoder(w).Encode(response)
}
