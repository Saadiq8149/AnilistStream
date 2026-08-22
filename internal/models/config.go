package models

type Config struct {
	// TODO: In future add cloudflare worker based stream proxying
	// CfWorkerURL      string `json:"cfWorkerURL"`
	AnilistAuthToken string `json:"anilistAuthToken,omitempty"`
	Provider         string `json:"provider"`
}
