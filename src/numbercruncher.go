package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func numbercruncher() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: go run . <fichier_entree> <fichier_sortie>")
		os.Exit(1)
	}

	inputFile := args[0]
	outputFile := args[1]

	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erreur: impossible de lire le fichier '%s': %v\n", inputFile, err)
		os.Exit(1)
	}

	lines := strings.Split(string(content), "\n")
	var resultLines []string

	for _, line := range lines {
		cleaned := cleanSpaces(line)

		processed := processLevel1(cleaned)

		resultLines = append(resultLines, processed)
	}

	outputContent := strings.Join(resultLines, "\n")
	err = os.WriteFile(outputFile, []byte(outputContent), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erreur lors de l'écriture dans '%s': %v\n", outputFile, err)
		os.Exit(1)
	}
}

func cleanSpaces(s string) string {
	fields := strings.Fields(s)
	return strings.Join(fields, " ")
}

func processLevel1(line string) string {
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