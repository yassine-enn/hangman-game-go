package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"text/tabwriter"
)

type Color string

const (
	Reset                   = "\x1b[0000m"
	Bright                  = "\x1b[0001m"
	BlackText               = "\x1b[0030m"
	RedText                 = "\x1b[0031m"
	GreenText               = "\x1b[0032m"
	YellowText              = "\x1b[0033m"
	BlueText                = "\x1b[0034m"
	MagentaText             = "\x1b[0035m"
	CyanText                = "\x1b[0036m"
	WhiteText               = "\x1b[0037m"
	DefaultText             = "\x1b[0039m"
	BrightRedText           = "\x1b[1;31m"
	BrightGreenText         = "\x1b[1;32m"
	BrightYellowText        = "\x1b[1;33m"
	BrightBlueText          = "\x1b[1;34m"
	BrightMagentaText       = "\x1b[1;35m"
	BrightCyanText          = "\x1b[1;36m"
	BrightWhiteText         = "\x1b[1;37m"
	BlackBackground         = "\x1b[0040m"
	RedBackground           = "\x1b[0041m"
	GreenBackground         = "\x1b[0042m"
	YellowBackground        = "\x1b[0043m"
	BlueBackground          = "\x1b[0044m"
	MagentaBackground       = "\x1b[0045m"
	CyanBackground          = "\x1b[0046m"
	WhiteBackground         = "\x1b[0047m"
	BrightBlackBackground   = "\x1b[0100m"
	BrightRedBackground     = "\x1b[0101m"
	BrightGreenBackground   = "\x1b[0102m"
	BrightYellowBackground  = "\x1b[0103m"
	BrightBlueBackground    = "\x1b[0104m"
	BrightMagentaBackground = "\x1b[0105m"
	BrightCyanBackground    = "\x1b[0106m"
	BrightWhiteBackground   = "\x1b[0107m"
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
A chaque fois que vous vous trompé, José s'approche davantage de la mort.
Ne laisse pas José se pendre ! Bonne chance !`

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
