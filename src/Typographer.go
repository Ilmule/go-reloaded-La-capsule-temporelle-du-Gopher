package main

import (
	"unicode"
)

func processLevel3(line string) string {
	if line == "" {
		return ""
	}

	runes := []rune(line)
	var result []rune

	for i := 0; i < len(runes); i++ {
		r := runes[i]

		if isPunctuation(r) {
			if len(result) > 0 && result[len(result)-1] == ' ' {
				result = result[:len(result)-1]
			}

			result = append(result, r)
			for i+1 < len(runes) && isPunctuation(runes[i+1]) {
				i++
				result = append(result, runes[i])
			}

			if i+1 < len(runes) {
				nextChar := runes[i+1]
				isBetweenDigits := len(result) >= 2 && unicode.IsDigit(result[len(result)-2]) && unicode.IsDigit(nextChar) && r == '.'

				if !isBetweenDigits && nextChar != ' ' {
					result = append(result, ' ')
				}
			}
		} else {
			result = append(result, r)
		}
	}

	return cleanSpaces(string(result))
}

func isPunctuation(r rune) bool {
	return r == ',' || r == '.' || r == '!' || r == '?' || r == ':' || r == ';'
}