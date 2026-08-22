package models

type Anime struct {
	Id            string    `json:"id"`
	Title         string    `json:"title"`
	Description   string    `json:"description,omitempty"`
	CoverImage    string    `json:"coverImage,omitempty"`
	BannerImage   string    `json:"bannerImage,omitempty"`
	Year          int       `json:"year,omitempty"`
	Status        string    `json:"status,omitempty"`
	Genres        []string  `json:"genres,omitempty"`
	TotalEpisodes int       `json:"totalEpisodes,omitempty"`
	Episodes      []Episode `json:"episodes,omitempty"`
}

type Episode struct {
	Id        string `json:"id"`
	Number    int    `json:"number"`
	Title     string `json:"title,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`
	Released  string `json:"released,omitempty"`
}

type AnimeSource struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url"`
	// HardSub SoftSub Dub Unknown
	SourceType string     `json:"sourceType,omitempty"`
	Subtitles  []Subtitle `json:"subtitles,omitempty"`
}
