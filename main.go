package main

import (
    "bufio"
    "fmt"
    "math/rand"
    "os"
    "strconv"
    "strings"
    "time"
)

// Lire le fichier words.txt et retourner un mot aléatoire
func getRandomWord() string {
    file, err := os.Open("words.txt")
    if err != nil {
        fmt.Println("Erreur lors de l'ouverture du fichier:", err)
        os.Exit(1)
    }
    defer file.Close()

    var words []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        words = append(words, scanner.Text())
    }

    rand.Seed(time.Now().UnixNano())
    return words[rand.Intn(len(words))]
}

// Lire le fichier hangman.txt et stocker chaque étape du pendu
func loadHangmanStages() []string {
    file, err := os.Open("hangman.txt")
    if err != nil {
        fmt.Println("Erreur lors de l'ouverture du fichier hangman.txt:", err)
        os.Exit(1)
    }
    defer file.Close()

    var stages []string
    scanner := bufio.NewScanner(file)
    var currentStage strings.Builder

    for scanner.Scan() {
        line := scanner.Text()
        if strings.TrimSpace(line) == "" || isNumeric(line) {
            if currentStage.Len() > 0 {
                stages = append(stages, currentStage.String())
                currentStage.Reset()
            }
        } else {
            currentStage.WriteString(line + "\n")
        }
    }
    if currentStage.Len() > 0 {
        stages = append(stages, currentStage.String())
    }
    return stages
}

// Vérifie si une chaîne est un nombre
func isNumeric(s string) bool {
    _, err := strconv.Atoi(s)
    return err == nil
}

// Afficher l'état actuel du pendu en fonction des erreurs
func printHangman(stages []string, errors int) {
    fmt.Println(stages[errors])
}

// Fonction principale pour jouer au pendu
func playGame() {
    word := getRandomWord()
    stages := loadHangmanStages()
    guessed := make([]rune, len(word))
    for i := range guessed {
        guessed[i] = '_'
    }

    attempts := 0
    maxAttempts := len(stages) - 1
    var incorrectLetters []rune

    for attempts < maxAttempts {
        fmt.Println("Mot actuel :", string(guessed))
        fmt.Println("Lettres incorrectes :", string(incorrectLetters))  // Afficher les lettres incorrectes
        fmt.Print("Devinez une lettre : ")
        var letter rune
        fmt.Scanf("%c\n", &letter)

        if !guessLetter(word, guessed, letter) {
            if !isInSlice(incorrectLetters, letter) {  // Vérifie que la lettre n'a pas déjà été devinée
                incorrectLetters = append(incorrectLetters, letter)
                attempts++
            } else {
                fmt.Println("Vous avez déjà deviné cette lettre incorrecte.")
            }
        }

        printHangman(stages, attempts)

        if string(guessed) == word {
            fmt.Println("Bien joué vous avez gagné ! Le mot était :", word)  // Affiche le mot complet
            return
        }
    }

    fmt.Println("Vous avez perdu. Le mot était :", word)  // Affiche le mot complet en cas de défaite
}

// Vérifie si une lettre est dans un slice de lettres
func isInSlice(slice []rune, letter rune) bool {
    for _, l := range slice {
        if l == letter {
            return true
        }
    }
    return false
}

// Met à jour les lettres devinées
func guessLetter(word string, guessed []rune, letter rune) bool {
    correct := false
    for i, char := range word {
        if char == letter {
            guessed[i] = letter
            correct = true
        }
    }
    return correct
}

func main() {
    playGame()
}
