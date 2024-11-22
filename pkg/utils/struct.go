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
	TestedLetters []string `json:"tested_letters"`
}

type Saves struct {
	Saves []Save `json:"saves"`
}
