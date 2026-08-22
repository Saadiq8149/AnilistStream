package models

type CatalogResponse struct {
	Metas []MetaPreview `json:"metas"`
}

type MetaPreview struct {
	Id          string   `json:"id"`
	Type        string   `json:"type"`
	Name        string   `json:"name"`
	Poster      string   `json:"poster"`
	Genres      []string `json:"genres,omitempty"`
	ImdbRating  string   `json:"imdbRating,omitempty"`
	ReleaseInfo string   `json:"releaseInfo,omitempty"`
	Description string   `json:"description,omitempty"`
}

type MetaResponse struct {
	Meta Meta `json:"meta"`
}

type Meta struct {
	Id          string   `json:"id"`
	Type        string   `json:"type"`
	Name        string   `json:"name"`
	Genres      []string `json:"genres,omitempty"`
	Poster      string   `json:"poster,omitempty"`
	Background  string   `json:"background,omitempty"`
	Description string   `json:"description,omitempty"`
	ReleaseInfo string   `json:"releaseInfo,omitempty"`
	ImdbRating  string   `json:"imdbRating,omitempty"`
	Videos      []Video  `json:"videos,omitempty"`
}

type Video struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Released  string `json:"released"`
	Episode   int    `json:"episode,omitempty"`
	Season    int    `json:"season,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`
}

type StreamResponse struct {
	Streams []Stream `json:"streams"`
}

type Stream struct {
	Name          string         `json:"name,omitempty"`
	Title         string         `json:"title,omitempty"`
	Url           string         `json:"url,omitempty"`
	Description   string         `json:"description,omitempty"`
	Subtitles     []Subtitle     `json:"subtitles,omitempty"`
	BehaviorHints *BehaviorHints `json:"behaviorHints,omitempty"`
}

type BehaviorHints struct {
	NotWebReady  bool          `json:"notWebReady,omitempty"`
	BingeGroup   string        `json:"bingeGroup,omitempty"`
	ProxyHeaders *ProxyHeaders `json:"proxyHeaders,omitempty"`
}

type ProxyHeaders struct {
	Request  map[string]string `json:"request,omitempty"`
	Response map[string]string `json:"response,omitempty"`
}

type Subtitle struct {
	Id   string `json:"id"`
	Lang string `json:"lang"`
	Url  string `json:"url"`
}
