package game

import (
	"bufio"
	"os"
	"strings"
)

func GetAsciiArt() ([]string, error) {
	file, err := os.Open("data/hangman.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var positions []string
	var asciiArt strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			positions = append(positions, asciiArt.String())
			asciiArt.Reset()
		} else {
			asciiArt.WriteString(line + "\n")
		}
	}
	if asciiArt.Len() > 0 {
		positions = append(positions, asciiArt.String())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return positions, nil
}

var currentPosition int
var positions []string

func InitPositions() error {
	var err error
	positions, err = GetAsciiArt()
	if err != nil {
		return err
	}
	currentPosition = 1
	return nil
}

func GetNextPosition() string {
	if currentPosition >= len(positions) {
		currentPosition = 0
	}
	position := positions[currentPosition]
	currentPosition++
	return position
}
