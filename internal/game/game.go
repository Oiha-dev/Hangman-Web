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

func ImportWords() []string {
	/*
		This function is used to import the words stored in the file data/words.txt
		Return: a slice of strings containing the words
	*/
	file, err := os.Open("data/words.txt")
	if err != nil {
		fmt.Errorf("Error: ", err)
		return nil
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Errorf("Error: ", err)
	}

	return words
}

func GetRandomWord(words []string) string {
	/*
		This function is used to get a random word from the slice of words
		Return: a string containing the word
	*/
	rand.Seed(time.Now().UnixNano())
	return words[rand.Intn(len(words))]
}

func RoundLogic(Jose *structure.HangManData, guessLetter string) {
	/*
		This function is used to update Jose with the guessed letter
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
