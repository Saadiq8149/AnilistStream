package types

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
	NotWebReady  bool                         `json:"notWebReady,omitempty"`
	BingeGroup   string                       `json:"bingeGroup,omitempty"`
	ProxyHeaders map[string]map[string]string `json:"proxyHeaders,omitempty"`
}
