package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

func main() {
	clearScreen()
	var x string
	posByte, err1 := ioutil.ReadFile("hangman.txt")
	if err1 != nil {
		log.Fatal(err1)
	}
	positions := strings.Split(string(posByte), "`,")
	var pos int = 0
	var replay string
	rand.Seed(time.Now().UnixNano())
	content, err := ioutil.ReadFile("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	wordsL := strings.Split(string(content), "\n") //list of words from words.txt
	min := 0
	max := len(wordsL)
	randomNumber := rand.Intn(max-min) + min //picks a random number between 0 and the number of lines of words.txt*
	word := wordsL[randomNumber]
	remainingAttempts := 10
	blankWord := DisplayBlankWord(word)
	var LetterAlreadyEntered []string
	var guessedWord string = DiplayRandLetters(blankWord, word)
	fmt.Println("Le mot à deviner est: ", guessedWord)
	fmt.Printf("Saisissez la lettre voulue > ")
	for remainingAttempts >= 0 && guessedWord != word {
		fmt.Scan(&x)
		LetterAlreadyEntered = append(LetterAlreadyEntered, x)
		fmt.Println("Les lettres que vous avez déjà saisies sont", LetterAlreadyEntered)
		isValid := CompareLetter(x, word)
		guessedWord = revealHiddenLetter(word, guessedWord, x, isValid)
		fmt.Println(guessedWord)
		if x == "quitter" {
			break
		}
		if !isValid {
			pos++
			remainingAttempts--
			fmt.Println(positions[pos])
			fmt.Println("Vous avez", remainingAttempts, "tentative(s) restante(s)")
			fmt.Printf("Saisissez la lettre voulue > ")
		} else {
			fmt.Println(positions[pos])
			fmt.Println("Vous avez", remainingAttempts, "tentative(s) restante(s)")
			if remainingAttempts != 1 {
				fmt.Printf("Saisissez la lettre voulue > ")
			}
		}
		if strings.Compare(strings.Join(strings.Split(guessedWord, " "), ""), word) == 0 && remainingAttempts > 0 {
			fmt.Println("Bien joué! Vous avez sauvé José!")
			break
		} else if remainingAttempts == 0 {
			fmt.Println("\n RIP José ", "\n Le mot à deviner était: ", word)
			fmt.Printf("Tapez oui pour recommencer une nouvelle partie > ")
			fmt.Scan(&replay)
			if replay == "oui" {
				fmt.Scan(&x)
				fmt.Printf("Saisissez la lettre voulue > ")
				remainingAttempts += 10
			}
		}
		if remainingAttempts <= 0 {
			break
		}
	}
}

func revealHiddenLetter(word string, guessWord string, InputLetter string, isInWord bool) string { //reveals the hidden letters
	guessWordL := strings.Split(guessWord, " ")
	fmt.Println(guessWordL)
	realWordL := strings.Split(word, "")
	for i := 0; i < len(guessWordL); i++ {
		if isInWord && InputLetter == realWordL[i] {
			guessWordL[i] = InputLetter
		}
	}
	return strings.Join(guessWordL, " ")
}

func DisplayBlankWord(randWord string) string { //displays the word in hidden letters
	randWordL := make([]string, len(randWord))
	randWordRune := []rune(randWord)
	for i := 0; i < len(randWordRune); i++ {
		if i == len(randWordRune)-1 {
			randWordL = append(randWordL, "_")
		} else {
			randWordL = append(randWordL, "_ ")
		}
	}

	return strings.Join(randWordL, "")
}

func CompareLetter(InputLetter string, Word string) bool { //checks whether the inputted letter is in in the word
	tabstring := []rune(Word)
	InputLetterR := []rune(InputLetter)
	var result bool = false
	for i := 0; i < len(tabstring); i++ {
		if tabstring[i] == InputLetterR[0] {
			result = true
		}
	}
	return result
}

func DiplayRandLetters(blankword string, randword string) string {
	n := len(randword)/2 - 1
	randomIndexL := make([]int, n)
	for k := 0; k < len(randomIndexL); k++ {
		randomIndexL[k] = rand.Intn(n)
	}
	randwordL := []rune(randword)
	blankwordL := []rune(blankword)
	for i := 0; i < len(randwordL); i++ {
		for j := 0; j < len(randomIndexL); j++ {
			if i == randomIndexL[j] {
				blankwordL[i*2] = randwordL[i]
			}
		}
	}
	return string(blankwordL)
}
