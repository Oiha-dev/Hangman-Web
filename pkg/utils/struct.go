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
	AttemptsLeft  int      `json:"attempts"`
}

type Saves struct {
	Saves []Save `json:"saves"`
}
