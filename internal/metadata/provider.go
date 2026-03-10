package metadata

import (
	"anilist-stream/internal/types"
	"os"
)

type MetadataProvider interface {
	Name() string
	SearchAnime(query string) ([]types.Metadata, error)
	GetAnime(id string) (types.Metadata, error)
}

type MetadataService struct {
	Provider MetadataProvider
}

func NewMetadataService() *MetadataService {
	selectedProvider := os.Getenv("METADATA_PROVIDER")
	if selectedProvider == "" {
		selectedProvider = "ANILIST" // Default provider
	}

	switch selectedProvider {
	case "ANILIST":
		return &MetadataService{Provider: NewAnilistProvider()}
	case "ALL_ANIME":
		return &MetadataService{Provider: NewAllAnimeProvider()}
	default:
		panic("Unsupported metadata provider: " + selectedProvider)
	}
}
