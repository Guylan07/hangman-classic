package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)
func printHangmanAscii() {
	fmt.Println(`
    _    _                                         
   | |  | |                                        
   | |__| | __ _ _ __   __ _ _ __ ___   __ _ _ __  
   |  __  |/ _' | '_ \ / _' | '_ ' _ \ / _' | '_ \ 
   | |  | | (_| | | | | (_| | | | | | | (_| | | | |
   |_|  |_|\__,_|_| |_|\__, |_| |_| |_|\__,_|_| |_|
                        __/ |                     
                       |___/                      
      `)
}
func chooseDifficulty() string {
	reader := bufio.NewReader(os.Stdin)
	var difficulty string

	for {
		fmt.Println("Choisissez une difficulté : Facile, Médium ou Difficile")
		input, _ := reader.ReadString('\n')
		difficulty = strings.TrimSpace(strings.ToLower(input))

		if difficulty == "facile" || difficulty == "médium" || difficulty == "difficile" {
			break
		} else {
			fmt.Println("Entrée invalide. Veuillez choisir entre Facile, Médium ou Difficile.")
		}
	}
	return difficulty
}
func loadWords(difficulty string) []string {
	var fileName string

	switch difficulty {
	case "facile":
		fileName = "words1.txt"
	case "médium":
		fileName = "words2.txt"
	case "difficile":
		fileName = "words3.txt"
	}

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		os.Exit(1)
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
		os.Exit(1)
	}

	return words
}
func pickRandomWord(words []string) string {
	rand.Seed(time.Now().UnixNano())
	return words[rand.Intn(len(words))]
}

func giveHint(word string) string {
	rand.Seed(time.Now().UnixNano())
	hintIndex := rand.Intn(len(word))
	return string(word[hintIndex])
}

func printHangman(tries int) {
	states := []string{
		`        
         
         
         
         
         
=========`,
		`         
      |  
      |  
      |  
      |  
      |  
=========`,
		`  +---+  
      |  
      |  
      |  
      |  
      |  
=========`,
		`  +---+  
  |   |  
      |  
      |  
      |  
      |  
=========`,
		`  +---+  
  |   |  
  O   |  
      |  
      |  
      |  
=========`,
		`  +---+  
  |   |  
  O   |  
  |   |  
      |  
      |  
=========`,
		`  +---+  
  |   |  
  O   |  
 /|   |  
      |  
      |  
=========`,
		`  +---+  
  |   |  
  O   |  
 /|\  |  
      |  
      |  
=========`,
		`  +---+  
  |   |  
  O   |  
 /|\  |  
 /    |  
      |  
=========`,
		`  +---+  
  |   |  
  O   |  
 /|\  |  
 / \  |  
      |  
=========`,
	}

	fmt.Println(states[tries])
}

func isValidLetter(input string) bool {
	if len(input) != 1 {
		return false
	}
	r := rune(input[0])
	return unicode.IsLetter(r)
}

func isWordFound(word string, guessedLetters map[string]bool) bool {
	for _, letter := range word {
		if _, found := guessedLetters[strings.ToLower(string(letter))]; !found {
			return false
		}
	}
	return true
}

//Pour gérer les accents ect
func normalizeLetter(r rune) rune {
	switch r {
	case 'à', 'â', 'ä':
		return 'a'
	case 'é', 'è', 'ê', 'ë':
		return 'e'
	case 'î', 'ï':
		return 'i'
	case 'ô', 'ö':
		return 'o'
	case 'ù', 'û', 'ü':
		return 'u'
	case 'ç':
		return 'c'
	default:
		return r
	}
}

func main() {
	printHangmanAscii()
	difficulty := chooseDifficulty()
	words := loadWords(difficulty)
	chosenWord := pickRandomWord(words)
	hint := giveHint(chosenWord)
	normalizedWord := strings.Map(normalizeLetter, strings.ToLower(chosenWord))
	guessedLetters := make(map[string]bool)
	wrongGuesses := []string{}
	lives := 10

	fmt.Printf("Indice : Le mot contient la lettre '%s'.\n", strings.ToLower(hint))
	guessedLetters[strings.ToLower(hint)] = true

	for lives > 0 {
		fmt.Print("Mot : ")
		for _, letter := range normalizedWord {
			if guessedLetters[string(letter)] {
				fmt.Printf("%c ", letter)
			} else {
				fmt.Print("_ ")
			}
		}
		fmt.Println()
		if len(wrongGuesses) > 0 {
			fmt.Println("Lettres incorrectes : ", strings.Join(wrongGuesses, ", "))
		}
		fmt.Printf("Vies restantes ♥︎ : %d\n", lives)
		printHangman(10 - lives)

		fmt.Println("Proposez une lettre ou devinez le mot complet :")

		reader := bufio.NewReader(os.Stdin)

		guess, _ := reader.ReadString('\n')

		guess = strings.TrimSpace(strings.ToLower(guess))

		guess = strings.Map(normalizeLetter, guess)

		if len(guess) > 1 {
			if guess == normalizedWord {
				fmt.Println("Bravo, vous avez deviné le mot complet ! Vous avez gagné !")
				return
			} else {
				fmt.Println("Mot incorrect ! Vous perdez 2 vies.")
				lives -= 2
			}
		} else {
			if !isValidLetter(guess) {
				fmt.Println("Veuillez entrer une lettre valide.")
				continue
			}

			if guessedLetters[guess] || contains(wrongGuesses, guess) {
				fmt.Println("Vous avez déjà essayé cette lettre.")
				continue
			}

			if strings.ContainsRune(normalizedWord, rune(guess[0])) {
				fmt.Printf("Bonne lettre : '%s' !\n", guess)
				guessedLetters[guess] = true
			} else {
				fmt.Printf("Mauvaise lettre : '%s'. Vous perdez une vie.\n", guess)
				wrongGuesses = append(wrongGuesses, guess)
				lives--
			}
		}
		if isWordFound(normalizedWord, guessedLetters) {
			fmt.Printf("Bravo, vous avez trouvé le mot : %s ! Vous avez gagné !\n", chosenWord)
			return
		}
	}

	fmt.Printf("Dommage, vous avez perdu ! Le mot était : %s.\n", chosenWord)
}
func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}


