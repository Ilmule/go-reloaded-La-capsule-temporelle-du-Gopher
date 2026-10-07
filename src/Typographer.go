package main

import (
	"unicode"
)

// Traite une ligne pour le Niveau 3 : Typographer (gestion de la ponctuation)
func processLevel3(line string) string {
	if line == "" {
		return ""
	}

	runes := []rune(line)
	var result []rune

	for i := 0; i < len(runes); i++ {
		r := runes[i]

		// 1. Gestion des signes de ponctuation : , . ! ? : ;
		if isPunctuation(r) {
			// Supprimer l'espace précédent s'il y en a un dans le résultat
			if len(result) > 0 && result[len(result)-1] == ' ' {
				result = result[:len(result)-1]
			}

			// Regrouper la ponctuation multiple (ex: ..., !?, !!)
			result = append(result, r)
			for i+1 < len(runes) && isPunctuation(runes[i+1]) {
				i++
				result = append(result, runes[i])
			}

			// Ajouter un espace après la ponctuation (sauf si c'est la fin de la ligne ou entre deux chiffres)
			if i+1 < len(runes) {
				nextChar := runes[i+1]
				// Vérifier si on est entre deux chiffres (ex: 3.14)
				isBetweenDigits := len(result) >= 2 && unicode.IsDigit(result[len(result)-2]) && unicode.IsDigit(nextChar) && r == '.'

				if !isBetweenDigits && nextChar != ' ' {
					result = append(result, ' ')
				}
			}
		} else {
			result = append(result, r)
		}
	}

	// Nettoyer les doubles espaces potentiels créés par les manipulations
	return cleanSpaces(string(result))
}

// Fonction auxiliaire pour identifier les ponctuations ciblées
func isPunctuation(r rune) bool {
	return r == ',' || r == '.' || r == '!' || r == '?' || r == ':' || r == ';'
}