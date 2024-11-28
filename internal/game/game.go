package game

import (
	"bufio"
	"fmt"
	classic_utils "hangman-web/pkg/hangman-classic/pkg/utils"
	"hangman-web/pkg/hangman-classic/structure"
	"math/rand"
	"os"
	"strings"
	"time"
)

func ImportWords() ([]string, []string, []string) {
	/*
		This function is used to import the words stored in the file data/words.txt
		Return: three slices of strings containing the words for easy, medium, and hard difficulties
	*/
	file, err := os.Open("data/words.txt")
	if err != nil {
		fmt.Errorf("Error: ", err)
		return nil, nil, nil
	}
	defer file.Close()

	var easyWords, mediumWords, hardWords []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		if len(word) < 6 {
			easyWords = append(easyWords, word)
		} else if len(word) < 10 {
			mediumWords = append(mediumWords, word)
		} else {
			hardWords = append(hardWords, word)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Errorf("Error: ", err)
	}

	return easyWords, mediumWords, hardWords
}

func GetRandomWord(easyWords, mediumWords, hardWords []string, difficulty string) string {
	/*
		This function is used to get a random word from the words slice based on the difficulty
		params: the words slices for easy, medium, and hard difficulties, the difficulty
		Return: a random word based on the difficulty
	*/
	rand.Seed(time.Now().UnixNano())
	switch difficulty {
	case "easy":
		return easyWords[rand.Intn(len(easyWords))]
	case "medium":
		return mediumWords[rand.Intn(len(mediumWords))]
	case "hard":
		return hardWords[rand.Intn(len(hardWords))]
	default:
		fmt.Println("Invalid difficulty")
		return ""
	}
}

func RoundLogic(Jose *structure.HangManData, guessLetter string) {
	/*
		This function is used to update Jose with the guessed letter
		params: the game data, the guessed letter
	*/
	Jose.GuessedLetters = append(Jose.GuessedLetters, strings.ToUpper(guessLetter))
	if !classic_utils.IsLetterInWord(Jose.Word, guessLetter) {
		Jose.HangmanState++
		Jose.Attempts--
	} else {
		classic_utils.UpdateWord(Jose, guessLetter)
		Jose.Score++
	}
}
