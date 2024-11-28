package classic_utils

import (
	"hangman-web/pkg/hangman-classic/structure"
	"math/rand"
)

func IsLetterInWord(word string, letter string) bool {
	/*
		This function is used to check if a letter is in a word
		Return: true if the letter is in the word, false otherwise
	*/
	for _, l := range word {
		if string(l) == letter {
			return true
		}
	}
	return false
}

func FirstPrintWord(word string) string {
	/*
		This function is used to print the first word with only half of the letters displayed
		Return: the word with half of the letters displayed
	*/
	lettersToDisplay := len(word)/2 - 1
	indexToDisplay := make([]int, lettersToDisplay)
	newWord := ""
	for i := 0; i < lettersToDisplay; i++ {
		randomIndex := rand.Intn(len(word))
		if i == 0 || (i > 0 && randomIndex != indexToDisplay[i-1]) {
			indexToDisplay[i] = rand.Intn(len(word))
		}
	}
	for i := 0; i < len(word); i++ {
		if containsInt(indexToDisplay, i) {
			newWord += string(word[i])
		} else {
			newWord += "_"
		}
	}
	return newWord
}

func UpdateWord(Jose *structure.HangManData, guessLetter string) {
	/*
		This function is used to update the word with the guessed letter
	*/
	newWord := ""
	for i := 0; i < len(Jose.Word); i++ {
		if string(Jose.Word[i]) == guessLetter {
			newWord += guessLetter
		} else {
			newWord += string(Jose.ToFind[i])
		}
	}
	Jose.ToFind = newWord
}

func ContainsStr(letters []string, s string) bool {
	/*
		This function is used to check if a string is in a slice of strings
		Return: true if the string is in the slice, false otherwise
	*/
	for _, l := range letters {
		if l == s {
			return true
		}
	}
	return false
}

func containsInt(indexes []int, i int) bool {
	/*
		This function is used to check if an integer is in a slice of integers
		Return: true if the integer is in the slice, false otherwise
	*/
	for _, index := range indexes {
		if index == i {
			return true
		}
	}
	return false
}

func GetPlayerRatio(player structure.Player) float64 {
	return float64(player.Score) / float64(player.Attempts)
}

func GetWinner(player1 structure.Player, player2 structure.Player) structure.Player {
	player1Ratio := GetPlayerRatio(player1)
	player2Ratio := GetPlayerRatio(player2)
	if player1Ratio > player2Ratio {
		return player1
	} else {
		return player2
	}
}
