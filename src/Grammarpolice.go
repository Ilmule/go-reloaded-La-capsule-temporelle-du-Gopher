package main

import (
	"strings"
	"unicode"
)

// Traite une ligne pour le Niveau 5 : Grammar Police (a -> an devant une voyelle)
func processLevel5(line string) string {
	if line == "" {
		return ""
	}

	tokens := strings.Split(line, " ")
	var result []string

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		// Vérifier si le token courant est un article indéfini ('a', 'A', 'an', 'An')
		if isArticle(token) {
			// Regarder s'il y a un élément suivant
			if i+1 < len(tokens) {
				next := tokens[i+1]
				
				// Extraire la première lettre en ignorant les apostrophes ou guillemets au début (ex: 'epic -> epic)
				firstLetter := getFirstLetter(next)

				if isVowel(firstLetter) {
					// Doit devenir 'an' ou 'An' selon la casse de l'article d'origine
					if token == "A" {
						result = append(result, "An")
					} else if token == "a" {
						result = append(result, "an")
					} else {
						// Si c'est déjà 'an' ou 'An', on le garde
						result = append(result, token)
					}
				} else {
					// Doit devenir 'a' ou 'A' si ce n'est pas une voyelle
					if token == "An" {
						result = append(result, "A")
					} else if token == "an" {
						result = append(result, "a")
					} else {
						// Si c'est déjà 'a' ou 'A', on le garde
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

// Vérifie si un token est un article indéfini a, A, an, An
func isArticle(s string) bool {
	return s == "a" || s == "A" || s == "an" || s == "An"
}

// Récupère la première lettre d'un mot en ignorant les guillemets/apostrophes de début
func getFirstLetter(s string) rune {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return r
		}
	}
	return 0
}

// Vérifie si une lettre est une voyelle (majuscule ou minuscule)
func isVowel(r rune) bool {
	rLower := unicode.ToLower(r)
	return rLower == 'a' || rLower == 'e' || rLower == 'i' || rLower == 'o' || rLower == 'u'
}