package isogram

import "strings"

func IsIsogram(word string) bool {
	seen := make(map[rune]bool)

	replacer := strings.NewReplacer("-", "", " ", "")
	word = replacer.Replace(word)
	word = strings.ToLower(word)

	for _, ch := range word {
		if seen[ch] {
			return false
		}
		seen[ch] = true
	}
	return true
}
