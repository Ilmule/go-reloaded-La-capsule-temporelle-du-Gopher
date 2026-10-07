package main

import (
	"strconv"
	"strings"
)

func processLevel2(line string) string {
	if line == "" {
		return ""
	}

	tokens := strings.Split(line, " ")
	var result []string

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		switch token {
		case "(up)":
			if len(result) > 0 {
				lastIdx := len(result) - 1
				val, err := strconv.ParseInt(result[lastIdx], 16, 64)
				if err == nil {
					result[lastIdx] = strconv.FormatInt(val, 10)
				}
			}

		case "(low)":
			if len(result) > 0 {
				lastIdx := len(result) - 1
				val, err := strconv.ParseInt(result[lastIdx], 2, 64)
				if err == nil {
					result[lastIdx] = strconv.FormatInt(val, 10)
				}
			}
		case "(cap)":
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