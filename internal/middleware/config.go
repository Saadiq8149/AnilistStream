package middleware

import (
	"aniliststream/internal/models"
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"
)

func ConfigMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		configString := r.PathValue("config")

		decoded, err := base64.RawURLEncoding.DecodeString(configString)
		if err != nil {
			http.Error(w, "invalid config", http.StatusBadRequest)
			return
		}

		var config models.Config
		if err := json.Unmarshal(decoded, &config); err != nil {
			http.Error(w, "invalid config", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), models.ConfigKey, config)
		next(w, r.WithContext(ctx))
	}
}

func GetConfig(r *http.Request) (models.Config, bool) {
	config, exists := r.Context().Value(models.ConfigKey).(models.Config)
	return config, exists
}
