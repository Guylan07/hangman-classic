package main

import (
	"fmt"
	"hangman/game"
)

func main() {
	fmt.Println("HANGMAN") // Affichage du titre en ASCII, éventuellement plus élaboré avec une librairie ou dessin
	game.Start()
}
