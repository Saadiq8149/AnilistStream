package types

type CatalogResponse struct {
	Metas []MetaPreview `json:"metas"`
}

type MetaPreview struct {
	ID          string   `json:"id"`
	Type        string   `json:"type"`
	Name        string   `json:"name"`
	Poster      string   `json:"poster"`
	Genres      []string `json:"genres,omitempty"`
	IMDBRating  string   `json:"imdbRating,omitempty"`
	ReleaseInfo string   `json:"releaseInfo,omitempty"`
	Description string   `json:"description,omitempty"`
}
