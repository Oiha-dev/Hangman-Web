package back

import (
	"encoding/json"
	"fmt"
	"hangman-web/pkg/utils"
	"io/ioutil"
	"net/http"
	"os"
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

	var scoreboardPlayer utils.Scoreboard

	for _, save := range saves.Saves {
		found := false
		for i, player := range scoreboardPlayer.Players {
			if player.Name == save.Username {
				score := scoreboardPlayer.Players[i].Score
				for _, letter := range save.TestedLetters {
					if utils.Contains(save.CurrentWord, letter) {
						score++
					}
				}
				scoreboardPlayer.Players[i].Score = score
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
			scoreboardPlayer.Players = append(scoreboardPlayer.Players, utils.Player{
				Name:     save.Username,
				Score:    score,
				Position: 1,
			})
		}
	}

	utils.SortPlayersByScore(scoreboardPlayer.Players)

	for i := range scoreboardPlayer.Players {
		scoreboardPlayer.Players[i].Position = i + 1
	}

	renderTemplate(w, "scoreboard/index", scoreboardPlayer)
}
