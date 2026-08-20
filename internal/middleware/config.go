package middleware

import (
	"aniliststream/internal/models"
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"
)

type contextKey string

const configKey contextKey = "config"

func ConfigMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")

		if len(parts) == 0 || parts[0] == "" {
			http.Error(w, "missing config", http.StatusBadRequest)
			return
		}

		decoded, err := base64.RawURLEncoding.DecodeString(parts[0])
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