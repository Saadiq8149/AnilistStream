package middleware

import (
	"net/http"
)

func StremioMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// TODO: Caching, How long to cache for?
		// w.Header().Set("Cache-Control", "max-age=3600")
		next.ServeHTTP(w, r)
	}
}
