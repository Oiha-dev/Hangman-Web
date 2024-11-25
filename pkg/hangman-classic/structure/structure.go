package structure

type HangManData struct {
	Word           string   // Word composed of '_', ex: H_ll_
	ToFind         string   // Final word chosen by the program at the beginning. It is the word to find
	Attempts       int      // Number of attempts left
	HangmanState   int      // It can be the array where the positions parsed in "hangman.txt" are stored
	GuessedLetters []string // Array of guessed letters
}

type Player struct {
	Name     string
	Score    int
	Attempts int
}
