package util

import "regexp"

var htmlRegex = regexp.MustCompile(`<.*?>`)

func StripHTML(s string) string {
	return htmlRegex.ReplaceAllString(s, "")
}
