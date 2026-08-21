package providers

import (
	"aniliststream/internal/models"
	"fmt"
)

type Manager struct {
	registry *Registry
}

func NewManager(registry *Registry) *Manager {
	return &Manager{
		registry: registry,
	}
}

func (m *Manager) GetProvider(name string) (Provider, bool) {
	provider, exists := m.registry.providers[name]
	return provider, exists
}

func (m *Manager) Search(providerName string, query string) ([]models.Anime, error) {
	if provider, exists := m.GetProvider(providerName); exists == true {
		return provider.Search(query)
	}
	return nil, fmt.Errorf("provider not found: %s", providerName)
}

func (m *Manager) GetDetails(providerName string, id string) (models.Anime, error) {
	if provider, exists := m.GetProvider(providerName); exists == true {
		return provider.GetDetails(id)
	}
	return models.Anime{}, fmt.Errorf("provider not found: %s", providerName)
}

func (m *Manager) GetStreams(providerName string, id string, episodeId string) ([]models.AnimeSource, error) {
	if provider, exists := m.GetProvider(providerName); exists == true {
		return provider.GetSources(id, episodeId)
	}
	return nil, fmt.Errorf("provider not found: %s", providerName)
}
