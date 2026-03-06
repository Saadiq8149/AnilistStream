package types

type Source struct {
	Name      string
	Url       string
	IsHLS     bool
	IsHardSub bool
	Subtitles []Subtitle
	Headers   map[string]string
}

type Subtitle struct {
	Id   string `json:"id"`
	Lang string `json:"lang"`
	Url  string `json:"url"`
}
