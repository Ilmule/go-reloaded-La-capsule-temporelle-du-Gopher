package main

import (
	"strings"
	"unicode"
)

func processLevel5(line string) string {
	if line == "" {
		return ""
	}

	tokens := strings.Split(line, " ")
	var result []string

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		if isArticle(token) {
			if i+1 < len(tokens) {
				next := tokens[i+1]
				
				firstLetter := getFirstLetter(next)

				if isVowel(firstLetter) {
					if token == "A" {
						result = append(result, "An")
					} else if token == "a" {
						result = append(result, "an")
					} else {
						result = append(result, token)
					}
				} else {
					if token == "An" {
						result = append(result, "A")
					} else if token == "an" {
						result = append(result, "a")
					} else {
						result = append(result, token)
					}
				}
				continue
			}
		}
		result = append(result, token)
	}

	return strings.Join(result, " ")
}

func isArticle(s string) bool {
	return s == "a" || s == "A" || s == "an" || s == "An"
}

func getFirstLetter(s string) rune {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return r
		}
	}
	return 0
}

func isVowel(r rune) bool {
	rLower := unicode.ToLower(r)
	return rLower == 'a' || rLower == 'e' || rLower == 'i' || rLower == 'o' || rLower == 'u'
}