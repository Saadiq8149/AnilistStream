package providers

import "aniliststream/internal/models"

type Provider interface {
	Search(query string) ([]models.Anime, error)
	GetDetails(id string) (models.Anime, error)
	GetSources(id string, episodeId string) ([]models.AnimeSource, error)
}
