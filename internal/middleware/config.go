package middleware

import (
	"aniliststream/internal/models"
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"
)

const configKey = "config"

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

		ctx := context.WithValue(r.Context(), configKey, config)
		next(w, r.WithContext(ctx))
	}
}
