package addon

import (
	"net/http"
	"os"
	"path/filepath"
)

func (h *Handler) Manifest(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/manifest.json")
}

func (h *Handler) Logo(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/logo.png")
}

func (h*Handler) Frontend(w http.ResponseWriter, r *http.Request) {
    dir := os.Getenv("FRONTEND_DIR")

    path := filepath.Join(dir, filepath.Clean(r.URL.Path))

    if info, err := os.Stat(path); err == nil && !info.IsDir() {
        http.ServeFile(w, r, path)
        return
    }

    http.ServeFile(w, r, filepath.Join(dir, "index.html"))
}
