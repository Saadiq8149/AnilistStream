package streams

import (
	"anilist-stream/internal/types"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	userAgent       = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/121.0"
	allanimeReferer = "https://allmanga.to"
	allanimeBase    = "https://allanime.day"
	allanimeAPI     = "https://api.allanime.day"
)

var allowedProviders = []string{"Default", "Yt-mp4", "S-mp4", "Luf-Mp4"}

var decoderMap = map[string]string{
	"79": "A", "7a": "B", "7b": "C", "7c": "D", "7d": "E", "7e": "F", "7f": "G",
	"70": "H", "71": "I", "72": "J", "73": "K", "74": "L", "75": "M", "76": "N", "77": "O",
	"68": "P", "69": "Q", "6a": "R", "6b": "S", "6c": "T", "6d": "U", "6e": "V", "6f": "W",
	"60": "X", "61": "Y", "62": "Z",
	"59": "a", "5a": "b", "5b": "c", "5c": "d", "5d": "e", "5e": "f", "5f": "g",
	"50": "h", "51": "i", "52": "j", "53": "k", "54": "l", "55": "m", "56": "n", "57": "o",
	"48": "p", "49": "q", "4a": "r", "4b": "s", "4c": "t", "4d": "u", "4e": "v", "4f": "w",
	"40": "x", "41": "y", "42": "z",
	"08": "0", "09": "1", "0a": "2", "0b": "3", "0c": "4", "0d": "5", "0e": "6", "0f": "7",
	"00": "8", "01": "9",
	"15": "-", "16": ".", "67": "_", "46": "~",
	"02": ":", "17": "/", "07": "?", "1b": "#",
	"63": "[", "65": "]", "78": "@", "19": "!",
	"1c": "$", "1e": "&", "10": "(", "11": ")",
	"12": "*", "13": "+", "14": ",", "03": ";",
	"05": "=", "1d": "%",
}

type AllAnimeProvider struct{}

type encodedProvider struct {
	Name      string
	EncodedID string
}

func NewAllAnimeProvider() *AllAnimeProvider {
	return &AllAnimeProvider{}
}

func (a *AllAnimeProvider) Name() string {
	return "AllAnime"
}

func (a *AllAnimeProvider) GetStreams(anilistID string, malID string, episode int) ([]types.Source, error) {
	// TODO: If anilistID missing we need to search by malID

	titles, err := getAnimeTitles(anilistID)
	if err != nil {
		return nil, fmt.Errorf("failed to get anime titles: %w", err)
	}

	var allAnimeID string
	for _, title := range titles {
		id, err := getAnimeByAnilistID(title, anilistID)
		if err == nil && id != "" {
			allAnimeID = id
			break
		}
	}

	if allAnimeID == "" {
		return nil, fmt.Errorf("anime not found on AllAnime")
	}

	providerIDs, err := fetchEncodedProviderIDs(allAnimeID, "sub", episode)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch provider IDs: %w", err)
	}

	dubProviderIDs, err := fetchEncodedProviderIDs(allAnimeID, "dub", episode)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch dub provider IDs: %w", err)
	}

	filteredProviders := []encodedProvider{}
	for _, provider := range providerIDs {
		for _, allowed := range allowedProviders {
			if provider.Name == allowed {
				filteredProviders = append(filteredProviders, provider)
				break
			}
		}
	}

	dubFilteredProviders := []encodedProvider{}
	for _, provider := range dubProviderIDs {
		for _, allowed := range allowedProviders {
			if provider.Name == allowed {
				dubFilteredProviders = append(dubFilteredProviders, provider)
				break
			}
		}
	}

	var sources []types.Source
	var mu sync.Mutex
	var wg sync.WaitGroup

	fetch := func(provider encodedProvider, label string) {
		defer wg.Done()

		decodedID := decodeID(provider.EncodedID)

		providerSources, err := fetchEpisodeSources(decodedID)
		if err != nil {
			return
		}

		for i := range providerSources {
			name := providerSources[i].Name

			if label == "Dub" {
				name = strings.ReplaceAll(name, "HardSub", "Dub")
				name = strings.ReplaceAll(name, "SoftSub", "Dub")
				name = strings.TrimSpace(name)
			}
			providerSources[i].Name = name
		}

		mu.Lock()
		sources = append(sources, providerSources...)
		mu.Unlock()
	}

	for _, provider := range filteredProviders {
		wg.Add(1)
		go fetch(provider, "Sub")
	}
	for _, provider := range dubFilteredProviders {
		wg.Add(1)
		go fetch(provider, "Dub")
	}

	wg.Wait()
	return sources, nil
}

func getAnimeTitles(anilistID string) ([]string, error) {
	query := `
	query ($id: Int) {
		Media(id: $id, type: ANIME) {
			title {
				romaji
				english
				native
			}
		}
	}`

	idInt, err := strconv.Atoi(anilistID)
	if err != nil {
		return nil, err
	}

	requestBody := map[string]any{
		"query": query,
		"variables": map[string]any{
			"id": idInt,
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

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Data struct {
			Media struct {
				Title struct {
					Romaji  string `json:"romaji"`
					English string `json:"english"`
					Native  string `json:"native"`
				} `json:"title"`
			} `json:"Media"`
		} `json:"data"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	titles := []string{}

	if result.Data.Media.Title.English != "" {
		titles = append(titles, result.Data.Media.Title.English)
	}

	if result.Data.Media.Title.Romaji != "" {
		titles = append(titles, result.Data.Media.Title.Romaji)
	}

	if result.Data.Media.Title.Native != "" {
		titles = append(titles, result.Data.Media.Title.Native)
	}

	return titles, nil
}

func getAnimeByAnilistID(title string, anilistID string) (string, error) {
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
					aniListId
				}
			}
		}`

	variables := map[string]any{
		"search": map[string]any{
			"allowAdult":   true,
			"allowUnknown": false,
			"query":        strings.ToLower(title),
		},
		"limit":           40,
		"translationType": "sub",
		"countryOrigin":   "ALL",
	}

	variablesJSON, err := json.Marshal(variables)
	if err != nil {
		return "", err
	}

	params := url.Values{}
	params.Set("query", searchGql)
	params.Set("variables", string(variablesJSON))

	reqURL := fmt.Sprintf("%s/api?%s", allanimeAPI, params.Encode())

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Referer", allanimeReferer)
	req.Header.Set("User-Agent", userAgent)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return "", err
	}

	var result struct {
		Data struct {
			Shows struct {
				Edges []struct {
					ID        string `json:"_id"`
					AniListID string `json:"aniListId"`
				} `json:"edges"`
			} `json:"shows"`
		} `json:"data"`
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	for _, edge := range result.Data.Shows.Edges {
		if edge.AniListID == anilistID {
			return edge.ID, nil
		}
	}

	return "", nil
}

func fetchEncodedProviderIDs(showID string, subOrDub string, episodeNum int) ([]encodedProvider, error) {
	query := `
	query (
		$showId: String!,
		$translationType: VaildTranslationTypeEnumType!,
		$episodeString: String!
	) {
		episode(
			showId: $showId
			translationType: $translationType
			episodeString: $episodeString
		) {
			episodeString
			sourceUrls
		}
	}`

	variables := map[string]any{
		"showId":          showID,
		"translationType": subOrDub,
		"episodeString":   strconv.Itoa(episodeNum),
	}

	variablesJSON, err := json.Marshal(variables)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	params.Set("query", query)
	params.Set("variables", string(variablesJSON))

	reqURL := fmt.Sprintf("%s/api?%s", allanimeAPI, params.Encode())

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, err
	}

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
			Episode struct {
				SourceUrls []struct {
					SourceName string `json:"sourceName"`
					SourceURL  string `json:"sourceUrl"`
				} `json:"sourceUrls"`
			} `json:"episode"`
		} `json:"data"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	providers := []encodedProvider{}
	for _, source := range result.Data.Episode.SourceUrls {
		encodedID := strings.TrimPrefix(source.SourceURL, "--")
		providers = append(providers, encodedProvider{
			Name:      source.SourceName,
			EncodedID: encodedID,
		})
	}

	return providers, nil
}

func decodeID(encodedID string) string {
	var splitChunks []string
	word := ""

	for _, char := range encodedID {
		word += string(char)
		if len(word) == 2 {
			splitChunks = append(splitChunks, word)
			word = ""
		}
	}

	decodedID := ""
	for _, chunk := range splitChunks {
		if decodedChar, ok := decoderMap[chunk]; ok {
			decodedID += decodedChar
		} else {
			decodedID += chunk
		}
	}

	decodedID = strings.Replace(decodedID, "clock", "clock.json", 1)
	return decodedID
}

func fetchEpisodeSources(providerID string) ([]types.Source, error) {
	var sources []types.Source

	if strings.HasPrefix(providerID, "https://tools.fast4speed.rsvp") {
		return sources, nil
	}

	reqURL := fmt.Sprintf("%s%s", allanimeBase, providerID)

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Referer", allanimeReferer)
	req.Header.Set("User-Agent", userAgent)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Links []struct {
			Link      string `json:"link"`
			HLS       bool   `json:"hls"`
			Subtitles []struct {
				Lang string `json:"lang"`
				URL  string `json:"src"`
			} `json:"subtitles"`
			Headers map[string]string `json:"headers"`
		} `json:"links"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	for _, link := range result.Links {
		videoLink := link.Link

		if link.HLS {
			var subtitles []types.Subtitle
			for _, sub := range link.Subtitles {
				subtitles = append(subtitles, types.Subtitle{
					Lang: sub.Lang,
					Url:  sub.URL,
				})
			}

			isHardSub := len(subtitles) == 0
			quality := "HLS"
			var softOrHard string
			if isHardSub {
				softOrHard = "HardSub"
			} else {
				softOrHard = "SoftSub"
			}
			sourceName := fmt.Sprintf("%s %s", quality, softOrHard)

			sources = append(sources, types.Source{
				Name:      sourceName,
				Url:       videoLink,
				IsHLS:     true,
				IsHardSub: isHardSub,
				Subtitles: subtitles,
			})
		} else {
			quality := "MP4 1080p"
			softOrHard := "HardSub"
			sourceName := fmt.Sprintf("%s %s", quality, softOrHard)

			sources = append(sources, types.Source{
				Name:      sourceName,
				Url:       videoLink,
				IsHLS:     false,
				IsHardSub: true,
				Subtitles: []types.Subtitle{},
			})
		}
	}

	return sources, nil
}
