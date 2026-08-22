package stremio

import "aniliststream/internal/providers"

type Handler struct {
	providerManager *providers.Manager
}

func NewHandler(providerManager *providers.Manager) *Handler {
	return &Handler{
		providerManager: providerManager,
	}
}
