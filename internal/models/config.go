package models

type Config struct {
	CFWorkerURL      string `json:"cfWorkerURL"`
	AnilistAuthToken string `json:"anilistAuthToken,omitempty"`
}