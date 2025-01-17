package livenshtein

import (
	"strings"
)

func Livenshtein(message1 string, message2 string) int {
	lignes := len(message1)
	colonnes := len(message2)
	var tableau = make([][]int, lignes+1)
	var tableaucout = make([][]int, lignes)
	for i := range tableau { // initialisation des matrices à 0
		tableau[i] = make([]int, colonnes+1)
		if i < lignes {
			tableaucout[i] = make([]int, colonnes)
		}
	}
	//ajout des indices allant de 0 à lignes le long des lignes, et les indices allant de 0 à colonnes le long des colonnes
	for i := range tableau {
		tableau[i][0] = i
	}
	for j := range tableau[0] {
		tableau[0][j] = j
	}

	for i := 0; i < lignes; i++ { //on calcule la matrice de coùt
		for j := 0; j < colonnes; j++ {
			if strings.EqualFold(string(message1[i]), string(message2[j])) { //on ignore les lettres capitales.
				tableaucout[i][j] = 0
			} else {
				tableaucout[i][j] = 1
			}
		}
	}

	for i := 1; i < lignes+1; i++ { //on calcule la matrice de substitution
		for j := 1; j < colonnes+1; j++ {
			tableau[i][j] = min(tableau[i-1][j]+1, tableau[i][j-1]+1, tableau[i-1][j-1]+tableaucout[i-1][j-1])
		}
	}
	res := tableau[lignes][colonnes]
	return res
}