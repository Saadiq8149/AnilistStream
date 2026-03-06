package pages

import (
	"html/template"
	"net/http"
	"os"
)

func ConfigureHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("./public/configure.html")

	tmpl.Execute(w, map[string]string{
		"ServerURL":   os.Getenv("SERVER_URL"),
		"ClientID":    os.Getenv("ANILIST_CLIENT_ID"),
		"RedirectURL": os.Getenv("ANILIST_REDIRECT_URL"),
	})
}
