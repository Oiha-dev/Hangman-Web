package utils

import (
	"bufio"
	"hangman-web/pkg/hangman-classic/structure"
	"os"
)

func SortPlayersByScore(players []Player) {
	for i := 0; i < len(players); i++ {
		for j := i + 1; j < len(players); j++ {
			if players[i].Score < players[j].Score {
				players[i], players[j] = players[j], players[i]
			}
		}
	}
}

func GetAsciiArt(position int) string {
	file, err := os.Open("data/hangman.txt")
	if err != nil {
		return "file not found"
	}
	defer file.Close()

	asciiArt := "\n"
	scanner := bufio.NewScanner(file)
	lineNumber := 1
	for scanner.Scan() {
		if lineNumber >= position*7-6 && lineNumber <= position*7 {
			asciiArt += scanner.Text() + "\n"
		}
		if lineNumber > position*7 {
			break
		}
		lineNumber++
	}
	return asciiArt
}

func isFinished(data structure.HangManData) bool {
	return data.Word == data.ToFind || data.Attempts == 0
}

func isWinned(data structure.HangManData) bool {
	return data.Word == data.ToFind
}
