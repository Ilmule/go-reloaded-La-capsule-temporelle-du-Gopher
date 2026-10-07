package main

import (
	"fmt"
	"os"
	"strings"
)

func lecteur() {
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
