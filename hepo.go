package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    // Ouvrir le fichier hangman.txt
    file, err := os.Open("hangman.txt")
    if err != nil {
        fmt.Println("Erreur lors de l'ouverture du fichier :", err)
        return
    }
    defer file.Close() // Fermer le fichier à la fin

    // Créer un lecteur pour le fichier
    scanner := bufio.NewScanner(file)

    // Lire le fichier ligne par ligne
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }

    // Vérifier les erreurs lors de la lecture
    if err := scanner.Err(); err != nil {
        fmt.Println("Erreur lors de la lecture du fichier :", err)
    }
}