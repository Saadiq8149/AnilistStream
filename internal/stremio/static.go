package stremio

import (
	"html/template"
	"net/http"
	"os"
)

type FrontendData struct {
	AnilistClientID string
}

func (h *Handler) Manifest(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/manifest.json")
}

func (h *Handler) Logo(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/logo.png")
}

func (h *Handler) Css(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	http.ServeFile(w, r, "public/style.css")
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("public/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FrontendData{
		AnilistClientID: os.Getenv("ANILIST_CLIENT_ID"),
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
