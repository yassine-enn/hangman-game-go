package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"text/tabwriter"
	"time"
)

type Color string

const (
	Reset                   Color = "\x1b[0000m"
	Bright                        = "\x1b[0001m"
	BlackText                     = "\x1b[0030m"
	RedText                       = "\x1b[0031m"
	GreenText                     = "\x1b[0032m"
	YellowText                    = "\x1b[0033m"
	BlueText                      = "\x1b[0034m"
	MagentaText                   = "\x1b[0035m"
	CyanText                      = "\x1b[0036m"
	WhiteText                     = "\x1b[0037m"
	DefaultText                   = "\x1b[0039m"
	BrightRedText                 = "\x1b[1;31m"
	BrightGreenText               = "\x1b[1;32m"
	BrightYellowText              = "\x1b[1;33m"
	BrightBlueText                = "\x1b[1;34m"
	BrightMagentaText             = "\x1b[1;35m"
	BrightCyanText                = "\x1b[1;36m"
	BrightWhiteText               = "\x1b[1;37m"
	BlackBackground               = "\x1b[0040m"
	RedBackground                 = "\x1b[0041m"
	GreenBackground               = "\x1b[0042m"
	YellowBackground              = "\x1b[0043m"
	BlueBackground                = "\x1b[0044m"
	MagentaBackground             = "\x1b[0045m"
	CyanBackground                = "\x1b[0046m"
	WhiteBackground               = "\x1b[0047m"
	BrightBlackBackground         = "\x1b[0100m"
	BrightRedBackground           = "\x1b[0101m"
	BrightGreenBackground         = "\x1b[0102m"
	BrightYellowBackground        = "\x1b[0103m"
	BrightBlueBackground          = "\x1b[0104m"
	BrightMagentaBackground       = "\x1b[0105m"
	BrightCyanBackground          = "\x1b[0106m"
	BrightWhiteBackground         = "\x1b[0107m"
)

const hangman = `                                                
 _   _    _    _   _  ____ __  __    _    _   _ 
| | | |  / \  | \ | |/ ___|  \/  |  / \  | \ | |
| |_| | / _ \ |  \| | |  _| |\/| | / _ \ |  \| |
|  _  |/ ___ \| |\  | |_| | |  | |/ ___ \| |\  |
|_| |_/_/   \_|_| \_|\____|_|  |_/_/   \_|_| \_|
`

const droits = `
by Juliette & Yassine
`

const rules = `
Regles : Vous avez 10 tentatives pour deviner le mot aléatoire.
A chaque fois que vous vous trompez, José s'approche davantage de la mort.
Ne laissez pas José se pendre ! Bonne chance !`

func main() {
	posByte, err1 := ioutil.ReadFile("hangman.txt")
	if err1 != nil {
		log.Fatal(err1)
	}
	positions := strings.Split(string(posByte), "`,")
	clearScreen()
	var pos int = 0
	var x string
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
	fmt.Printf("Saisissez la lettre voulue > ")
	var guessedWord string = DiplayRandLetters(DisplayBlankWord(word), word)
	for remainingAttempts >= 0 && guessedWord != word {
		fmt.Println(guessedWord)
		fmt.Scan(&x)
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

func DisplayBlankWord(randWord string) string { //displays the word in hidden letters
	randWordL := make([]string, len(randWord))
	for i := 0; i < len(randWord); i++ {
		if i == len(randWord)-1 {
			randWordL = append(randWordL, "_")
		} else {
			randWordL = append(randWordL, "_ ")
		}
	}

	return strings.Join(randWordL, "")
}

func DiplayRandLetters(blankword string, randword string) string {
	n := len(randword)/2 - 1
	randomIndexL := make([]int, n)
	rand.Seed(time.Now().UnixNano())
	for k := 0; k < n; k++ {
		randomIndexL[k] = rand.Intn(len(randword))
	}
	blankwordL := strings.Split(blankword, "")
	randwordL := strings.Split(randword, "")
	for i := 0; i < len(blankword); i++ {
		for j := 0; j < len(randomIndexL); j++ {
			if i == randomIndexL[j] {
				blankwordL[i*2] = randwordL[i]
			}
		}
	}
	return strings.Join(blankwordL, "")
}

func revealHiddenLetter(word string, guessWord string, InputLetter string, isInWord bool) string { //reveals the hidden letters
	guessWordL := strings.Split(guessWord, " ")
	realWordL := strings.Split(word, "")
	for i := 0; i < len(guessWordL); i++ {
		if isInWord && InputLetter == realWordL[i] && guessWordL[i] == "_" {
			guessWordL[i] = InputLetter
		}
	}
	return strings.Join(guessWordL, " ")
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

func clearScreen() { // permet d'effacer les affichages précédents sauf le logo HANGMAN
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0) //  gestion des éspaces autour des affichages afin d'avoir quelque chose d'aligner
	colors := []Color{BrightCyanText, BrightGreenText, RedBackground}
	PrintRow(writer, PaintRow(colors, []string{hangman, droits, rules}))
}

//---------------------------------------------------------------

// traitement couleur

func (c *Color) String() string {
	return fmt.Sprintf("%v", c)
}

func Paint(color Color, value string) string {
	return fmt.Sprintf("%v%v%v", color, value, Reset)
}

func PaintRow(colors []Color, row []string) []string {
	paintedRow := make([]string, len(row))
	for i, v := range row {
		paintedRow[i] = Paint(colors[i], v)
	}
	return paintedRow
}

func PrintRow(writer io.Writer, line []string) {
	fmt.Fprintln(writer, strings.Join(line, "\t"))
}

//--------------------------------------------------------
