package types

type MetaResponse struct {
	Meta Meta `json:"meta"`
}

type Meta struct {
	ID          string   `json:"id"`
	Type        string   `json:"type"`
	Name        string   `json:"name"`
	Genres      []string `json:"genres,omitempty"`
	Poster      string   `json:"poster,omitempty"`
	Background  string   `json:"background,omitempty"`
	Description string   `json:"description,omitempty"`
	ReleaseInfo string   `json:"releaseInfo,omitempty"`
	IMDBRating  string   `json:"imdbRating,omitempty"`
	Videos      []Video  `json:"videos,omitempty"`
}

type Video struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Released string `json:"released"`
	Episode  int    `json:"episode,omitempty"`
	Season   int    `json:"season,omitempty"`
}
