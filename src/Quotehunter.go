package main

import (
	"strings"
)

func processLevel4(line string) string {
	if line == "" {
		return ""
	}
	
	line = strings.ReplaceAll(line, "'", " ' ")

	tokens := strings.Fields(line)
	var result []string

	openQuote := true

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		switch token {
		case "'":
			if openQuote == true {
				if i+1 < len(tokens) {
					tokens[i+1] = "'" + tokens[i+1]
				} else {
					result = append(result, "'")
				}
				openQuote = false
			} else {
				if len(result) > 0 {
					dernierIndex := len(result) - 1
					result[dernierIndex] = result[dernierIndex] + "'"
				} else {
					result = append(result, "'")
				}
				openQuote = true
			}

		default:
			result = append(result, token)
		}
	}

	return strings.Join(result, " ")
}