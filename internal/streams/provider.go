package streams

import (
	"anilist-stream/internal/types"
	"fmt"
	"os"
	"strings"
)

type SourceProvider interface {
	Name() string
	GetStreams(anilistID string, malID string, episode int) ([]types.Source, error)
}

type SourceService struct {
	providers []SourceProvider
}

func NewSourceService() *SourceService {
	selectedProviders := strings.Split(os.Getenv("SOURCE_PROVIDERS"), ",")
	providers := []SourceProvider{}

	for _, providerName := range selectedProviders {
		switch providerName {
		case "ALL_ANIME":
			providers = append(providers, NewAllAnimeProvider())
		}
	}

	return &SourceService{
		providers: providers,
	}
}

func (s *SourceService) GetStreams(anilistID string, malID string, episode int) ([]types.Source, error) {
	var allSources []types.Source

	for _, provider := range s.providers {
		sources, err := provider.GetStreams(anilistID, malID, episode)
		if err != nil {
			continue
		}
		for i := range sources {
			sources[i].Name = fmt.Sprintf("AnilistStream %s %s", provider.Name(), sources[i].Name)
		}
		allSources = append(allSources, sources...)
	}

	return allSources, nil
}
