package src

import (
	"strconv"
	"strings"
)

func ProcessLevel1(line string) string {
	if line == "" {
		return ""
	}

	tokens := strings.Split(line, " ")
	var result []string

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		switch token {
		case "(hex)":
			if len(result) > 0 {
				lastIdx := len(result) - 1
				val, err := strconv.ParseInt(result[lastIdx], 16, 64)
				if err == nil {
					result[lastIdx] = strconv.FormatInt(val, 10)
				}
			}

		case "(bin)":
			if len(result) > 0 {
				lastIdx := len(result) - 1
				val, err := strconv.ParseInt(result[lastIdx], 2, 64)
				if err == nil {
					result[lastIdx] = strconv.FormatInt(val, 10)
				}
			}
		
		default:
			result = append(result, token)
		}
	}

	return strings.Join(result, " ")
}