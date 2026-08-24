package providers

import (
	"aniliststream/internal/models"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	neturl "net/url"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type AniDb struct{}

const BASE_API = "https://anidb.app"

func NewAniDbProvider() AniDb {
	return AniDb{}
}

func (a AniDb) Name() string {
	return "AniDb"
}

func (a AniDb) get(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36",
	)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (a AniDb) Search(query string) ([]models.AnimePreview, error) {
	url := BASE_API + "/browse?q=" + url.QueryEscape(query)
	body, err := a.get(url)
	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	var results []models.AnimePreview

	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if !exists {
			return
		}
		title, exists := s.Attr("title")
		if !exists {
			return
		}
		coverImage, exists := s.Find("img").Attr("src")
		if !exists {
			return
		}
		parsed, err := neturl.Parse(href)
		if err != nil {
			return
		}
		id := strings.TrimPrefix(parsed.Path, "/anime/")
		results = append(results, models.AnimePreview{
			Id:         id,
			Title:      title,
			CoverImage: coverImage,
		})
	})

	return results, nil
}

func (a AniDb) GetDetails(id string) (models.Anime, error) {
	rawURL := BASE_API + "/anime/" + id

	body, err := a.get(rawURL)
	if err != nil {
		return models.Anime{}, err
	}
	err = os.WriteFile("body.html", body, 0644)
	if err != nil {
		return models.Anime{}, err
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		return models.Anime{}, err
	}

	title := strings.TrimSpace(doc.Find("h1").First().Text())

	description := strings.TrimSpace(
		doc.Find("h2").FilterFunction(func(_ int, s *goquery.Selection) bool {
			return strings.TrimSpace(s.Text()) == "Synopsis"
		}).Parent().Find("p").First().Text(),
	)

	coverImage, _ := doc.Find(`img[alt="` + title + `"]`).First().Attr("src")

	status := ""
	doc.Find("dt").EachWithBreak(func(_ int, s *goquery.Selection) bool {
		if strings.TrimSpace(s.Text()) == "Status" {
			status = strings.TrimSpace(s.Next().Find("a").Text())
			return false
		}
		return true
	})

	year := 0
	doc.Find("dt").EachWithBreak(func(_ int, s *goquery.Selection) bool {
		if strings.TrimSpace(s.Text()) == "Season" {
			season := strings.TrimSpace(s.Next().Find("a").Text())
			parts := strings.Fields(season)
			if len(parts) == 2 {
				year, _ = strconv.Atoi(parts[1])
			}
			return false
		}
		return true
	})

	var genres []string
	doc.Find("h1").First().Parent().Parent().
		Find(`a[href^="/genres/"]`).
		Each(func(_ int, s *goquery.Selection) {
			genre := strings.TrimSpace(s.Text())
			if genre != "" {
				genres = append(genres, genre)
			}
		})

	idParts := strings.Split(id, "-")
	numericId := idParts[len(idParts)-1]

	episodeURL := BASE_API + "/api/frontend/anime/" + numericId + "/episodes"
	episodeBody, err := a.get(episodeURL)
	if err != nil {
		return models.Anime{}, err
	}
	type aniDbEpisode struct {
		ID     int  `json:"id"`
		Number int  `json:"number"`
		Filler bool `json:"filler"`
	}
	type aniDbEpisodesResponse struct {
		Episodes []aniDbEpisode `json:"episodes"`
	}

	var episodeResponse aniDbEpisodesResponse
	if err := json.Unmarshal(episodeBody, &episodeResponse); err != nil {
		return models.Anime{}, err
	}
	episodes := make([]models.Episode, 0, len(episodeResponse.Episodes))
	for _, episode := range episodeResponse.Episodes {
		episodes = append(episodes, models.Episode{
			Id:       strconv.Itoa(episode.ID),
			Number:   episode.Number,
			Title:    "Episode " + strconv.Itoa(episode.Number),
			Released: "1999-10-20T00:00:00.000Z",
		})
	}
	bannerImage, _ := doc.Find(`main > div.relative img`).First().Attr("src")
	rating := ""
	doc.Find("span.badge-gray").EachWithBreak(func(_ int, s *goquery.Selection) bool {
		text := strings.TrimSpace(s.Text())
		if text != "" {
			rating = text
			return false
		}
		return true
	})

	anime := models.Anime{
		Id:            id,
		Title:         title,
		Description:   description,
		CoverImage:    coverImage,
		Status:        status,
		Year:          year,
		Genres:        genres,
		Episodes:      episodes,
		TotalEpisodes: len(episodes),
		BannerImage:   bannerImage,
		Rating:        rating,
	}
	return anime, nil
}

func (a AniDb) GetSources(id string, episodeId string) ([]models.AnimeSource, error) {
	episodeURL := BASE_API + "/api/frontend/episode/" + episodeId + "/languages"
	body, err := a.get(episodeURL)
	if err != nil {
		return nil, err
	}

	type language struct {
		Code     string `json:"code"`
		Name     string `json:"name"`
		EmbedURL string `json:"embed_url"`
	}
	type languagesResponse struct {
		Languages []language `json:"languages"`
	}
	var languageResponse languagesResponse
	if err := json.Unmarshal(body, &languageResponse); err != nil {
		return nil, err
	}

	var embeds []language
	for _, language := range languageResponse.Languages {
		embeds = append(embeds, language)
	}
	var sources []models.AnimeSource

	for _, language := range languageResponse.Languages {
		body, err := a.get(language.EmbedURL)
		if err != nil {
			return nil, err
		}

		re := regexp.MustCompile(`file:\s*'([^']+)'`)
		match := re.FindSubmatch(body)
		if len(match) < 2 {
			continue
		}

		streamURL := string(match[1])

		sourceType := "Unknown"
		switch language.Code {
		case "eng":
			sourceType = "Dub"
		case "jpn":
			sourceType = "HardSub"
		}

		sources = append(sources, models.AnimeSource{
			Name:       a.Name(),
			Url:        streamURL,
			SourceType: sourceType,
		})
	}

	return sources, nil
}
