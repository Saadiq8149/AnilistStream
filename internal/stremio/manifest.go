package stremio

import (
	"net/http"
)

func ManifestHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/manifest.json")
}
