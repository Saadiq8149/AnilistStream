package middleware

import (
	"fmt"
	"net/http"
)

func StremioMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("route hit:", r.URL.Path)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// TODO: Caching, How long to cache for?
		// w.Header().Set("Cache-Control", "max-age=3600")
		next.ServeHTTP(w, r)
	}
}
