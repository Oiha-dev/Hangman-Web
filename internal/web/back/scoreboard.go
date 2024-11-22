package back

import (
	"encoding/json"
	"fmt"
	"hangman-web/pkg/utils"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
)

func scoreboard(w http.ResponseWriter, r *http.Request) {
	jsonFile, err := os.Open("data/save.json")
	if err != nil {
		fmt.Println("Failed to open JSON file:", err)
		return
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Failed to read JSON file:", err)
		return
	}

	var saves utils.Saves
	err = json.Unmarshal(byteValue, &saves)
	if err != nil {
		fmt.Println("Failed to unmarshal JSON data:", err)
		return
	}

	var scoreboard utils.Scoreboard

	// Calculate scores first
	for _, save := range saves.Saves {
		found := false
		for i, player := range scoreboard.Players {
			if player.Name == save.Username {
				score := scoreboard.Players[i].Score
				for _, letter := range save.TestedLetters {
					if utils.Contains(save.CurrentWord, letter) {
						score++
					}
				}
				scoreboard.Players[i].Score = score
				found = true
				break
			}
		}
		if !found {
			score := 0
			for _, letter := range save.TestedLetters {
				if utils.Contains(save.CurrentWord, letter) {
					score++
				}
			}
			scoreboard.Players = append(scoreboard.Players, utils.Player{
				Name:     save.Username,
				Score:    score,
				Position: 1, // Default position, will be updated next
			})
		}
	}

	// Sort players by score (highest to lowest)
	sort.Slice(scoreboard.Players, func(i, j int) bool {
		return scoreboard.Players[i].Score > scoreboard.Players[j].Score
	})

	// Assign positions (1-based)
	for i := range scoreboard.Players {
		scoreboard.Players[i].Position = i + 1
	}

	renderTemplate(w, "scoreboard/index", scoreboard)
}
