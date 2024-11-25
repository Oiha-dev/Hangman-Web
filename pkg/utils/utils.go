package utils

import "hangman-web/pkg/hangman-classic/structure"

func SortPlayersByScore(players []Player) {
	for i := 0; i < len(players); i++ {
		for j := i + 1; j < len(players); j++ {
			if players[i].Score < players[j].Score {
				players[i], players[j] = players[j], players[i]
			}
		}
	}
}

func isFinished(data structure.HangManData) bool {
	return data.Word == data.ToFind || data.Attempts == 0
}

func isWinned(data structure.HangManData) bool {
	return data.Word == data.ToFind
}
