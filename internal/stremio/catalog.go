package stremio

import (
	"fmt"
	"net/http"
	"strings"
)

func CatalogHandler(w http.ResponseWriter, r *http.Request) {
	params := strings.Split(r.URL.Path, "/")
	catalogID := params[3]

	if len(params) > 4 {
		extra := strings.TrimSuffix(params[4], ".json")

		searchQuery := strings.TrimPrefix(extra, "search=")
		fmt.Println(searchQuery)
	}

	fmt.Println(catalogID)
}
