package utils

func Contains(word string, letter string) bool {
	for _, char := range word {
		if string(char) == letter {
			return true
		}
	}
	return false
}

func SortPlayersByScore(players []Player) {
	for i := 0; i < len(players); i++ {
		for j := i + 1; j < len(players); j++ {
			if players[i].Score < players[j].Score {
				players[i], players[j] = players[j], players[i]
			}
		}
	}
}
