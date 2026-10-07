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

		if token == "(up," || token == "(low," || token == "(cap," {
			if i+1 < len(tokens) {
				nextToken := tokens[i+1]
				numberInString := nextToken[0 : len(nextToken)-1]
				count, err := strconv.Atoi(numberInString)

				if err == nil {
					startIndex := len(result) - count
					if startIndex < 0 {
						startIndex = 0
					}
					for j := startIndex; j < len(result); j++ {
						if token == "(up," {
							result[j] = strings.ToUpper(result[j])
						}
						if token == "(low," {
							result[j] = strings.ToLower(result[j])
						}
						if token == "(cap," {
							word := result[j]
							if len(word) > 0 {
								firstLetter := strings.ToUpper(string(word[0]))
								restOfWord := strings.ToLower(word[1:])
								result[j] = firstLetter + restOfWord
							}
						}
					}

					
					i++
					continue
				}
			}
		}

		switch token {
		case "(up)":
			if len(result) > 0 {
				lastIdx := len(result) - 1
				result[lastIdx] = strings.ToUpper(result[lastIdx])
			}

		case "(low)":
			if len(result) > 0 {
				lastIdx := len(result) - 1
				result[lastIdx] = strings.ToLower(result[lastIdx])
			}

		case "(cap)":
			if len(result) > 0 {
				lastIdx := len(result) - 1
				word := result[lastIdx]
				if len(word) > 0 {
					result[lastIdx] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
				}
			}

		default:
			result = append(result, token)
		}
	}

	return strings.Join(result, " ")
}