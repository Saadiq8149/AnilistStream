package models

type Config struct {
	CfWorkerURL      string `json:"cfWorkerURL"`
	AnilistAuthToken string `json:"anilistAuthToken,omitempty"`
	Provider         string `json:"provider"`
}
