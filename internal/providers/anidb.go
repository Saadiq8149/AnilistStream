package providers

import "aniliststream/internal/models"

type AniDb struct{}

func NewAniDbProvider() AniDb {
	return AniDb{}
}

func (a AniDb) Name() string {
	return "AniDb"
}

func (a AniDb) Search(query string) ([]models.Anime, error) {
	return nil, nil
}

func (a AniDb) GetDetails(id string) (models.Anime, error) {
	return models.Anime{}, nil
}

func (a AniDb) GetSources(id string, episodeId string) ([]models.AnimeSource, error) {
	return nil, nil
}
