package types

type Metadata struct {
	ProviderID  string
	AnilistID   string
	MalID       string
	Title       string
	Description string
	Poster      string
	Banner      string
	FromYear    int
	ToYear      int
	Rating      float64
	Episodes    int
	Genres      []string
}
