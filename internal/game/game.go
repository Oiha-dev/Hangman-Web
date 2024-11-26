package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
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
