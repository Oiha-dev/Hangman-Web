package main

import (
	"hangman-web/internal/web/back"
)

func main() {
	// word := game.GetRandomWord(game.ImportWords())
	// fmt.Println(word, classic_utils.FirstPrintWord(word))
	back.StartServer()
}
