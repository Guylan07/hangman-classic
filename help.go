package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Liste de mots possibles (ici, un seul mot choisi)
	mot := "golang"
	tentatives := 6
	lettresDevinees := make(map[rune]bool)

	// Créer un scanner pour lire l'entrée utilisateur
	scanner := bufio.NewScanner(os.Stdin)

	for tentatives > 0 {
		fmt.Printf("Tentatives restantes: %d\n", tentatives)
		afficherMot(mot, lettresDevinees)

		// Demander une lettre à l'utilisateur
		fmt.Print("Devinez une lettre: ")
		scanner.Scan()
		lettre := strings.ToLower(scanner.Text())

		// Vérifier que l'utilisateur a bien entré une seule lettre
		if len(lettre) != 1 {
			fmt.Println("Veuillez entrer une seule lettre.")
			continue
		}

		// Convertir la lettre en rune
		lettreRune := rune(lettre[0])

		// Si la lettre a déjà été devinée
		if lettresDevinees[lettreRune] {
			fmt.Println("Vous avez déjà deviné cette lettre.")
			continue
		}

		// Ajouter la lettre aux devinettes
		lettresDevinees[lettreRune] = true

		// Vérifier si la lettre fait partie du mot
		if strings.ContainsRune(mot, lettreRune) {
			fmt.Println("Bonne lettre !")
			if motDevine(mot, lettresDevinees) {
				fmt.Println("Félicitations, vous avez gagné ! Le mot était :", mot)
				return
			}
		} else {
			fmt.Println("Mauvaise lettre.")
			tentatives--
		}
	}

	fmt.Println("Vous avez perdu. Le mot était :", mot)
}

// Fonction pour afficher le mot partiellement découvert
func afficherMot(mot string, lettresDevinees map[rune]bool) {
	for _, lettre := range mot {
		if lettresDevinees[lettre] {
			fmt.Printf("%c ", lettre)
		} else {
			fmt.Print("_ ")
		}
	}
	fmt.Println()
}

// Fonction pour vérifier si le mot entier a été deviné
func motDevine(mot string, lettresDevinees map[rune]bool) bool {
	for _, lettre := range mot {
		if !lettresDevinees[lettre] {
			return false
		}
	}
	return true
}
