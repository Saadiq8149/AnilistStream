package providers

import (
	"aniliststream/internal/models"
	"fmt"
	"maps"
	"slices"
)

type Manager struct {
	registry *Registry
}

func NewManager() *Manager {
	return &Manager{
		registry: NewRegistry(),
	}
}

func (m *Manager) GetProvidersList() []string {
	return slices.Collect(maps.Keys(m.registry.providers))
}

func (m *Manager) GetProvider(name string) (Provider, bool) {
	provider, exists := m.registry.providers[name]
	return provider, exists
}

func (m *Manager) Search(providerName string, query string) ([]models.Anime, error) {
	if provider, exists := m.GetProvider(providerName); exists {
		return provider.Search(query)
	}
	return nil, fmt.Errorf("provider not found: %s", providerName)
}

func (m *Manager) GetDetails(providerName string, id string) (models.Anime, error) {
	if provider, exists := m.GetProvider(providerName); exists {
		return provider.GetDetails(id)
	}
	return models.Anime{}, fmt.Errorf("provider not found: %s", providerName)
}

func (m *Manager) GetStreams(providerName string, id string, episodeId string) ([]models.AnimeSource, error) {
	if provider, exists := m.GetProvider(providerName); exists {
		return provider.GetSources(id, episodeId)
	}
	return nil, fmt.Errorf("provider not found: %s", providerName)
}
