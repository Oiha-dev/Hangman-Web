package utils

type Player struct {
	Name     string
	Score    int
	Position int
}

type Scoreboard struct {
	Players []Player
}

type Save struct {
	Username      string   `json:"username"`
	CurrentWord   string   `json:"current_word"`
	GoalWord      string   `json:"goal_word"`
	TestedLetters []string `json:"tested_letters"`
	Score         int      `json:"score"`
	AttemptsLeft  int      `json:"attempts_left"`
}

type Saves struct {
	Saves []Save `json:"saves"`
}

type HangManDataMulti struct {
	Word           string   // Word composed of '_', ex: H_ll_
	ToFind         string   // Final word chosen by the program at the beginning. It is the word to find
	Attempts       int      // Number of attempts left
	HangmanState   int      // It can be the array where the positions parsed in "hangman.txt" are stored
	GuessedLetters []string // Array of guessed letters
	GuessedWords   []string // Array of guessed words
	Score          int      // Score of the player
	IsWinned       bool     // True if the player has won, false otherwise
	Player1        Player   // Player 1 data
	Player2        Player   // Player 2 data
	CurrentPlayer  int      // Current player
}
