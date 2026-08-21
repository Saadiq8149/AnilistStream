package providers

type Registry struct {
	providers map[string]Provider
}

func NewRegistry() *Registry {
	registry := &Registry{
		providers: make(map[string]Provider),
	}
	aniDbProvider := NewAniDbProvider()
	registry.providers[aniDbProvider.Name()] = aniDbProvider
	return registry
}
