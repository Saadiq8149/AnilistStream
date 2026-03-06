package pages

import (
	"html/template"
	"net/http"
	"os"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("./public/index.html")

	tmpl.Execute(w, map[string]string{
		"ServerURL": os.Getenv("SERVER_URL"),
	})
}
