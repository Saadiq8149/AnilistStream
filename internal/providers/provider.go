package providers

import "aniliststream/internal/models"

// future todo maybe have a central metadata provider and multiple stream providers
// with a id mapping service to map anime ids between providers so we can get a lot more
// sources and streams

type Provider interface {
	Search(query string) ([]models.AnimePreview, error)
	GetDetails(id string) (models.Anime, error)
	GetSources(id string, episodeId string) ([]models.AnimeSource, error)
}
