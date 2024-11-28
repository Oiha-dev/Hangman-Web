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
	/*
		This function is used to display the scoreboard of the game
		by getting the saves from the save.json file
	*/
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
		if utils.IsPlayerInScoreboard(utils.Player{Name: save.Username}, scoreboardPlayer) {
			for i := range scoreboardPlayer.Players {
				if scoreboardPlayer.Players[i].Name == save.Username {
					scoreboardPlayer.Players[i].Score += save.Score
					break
				}
			}
		} else {
			scoreboardPlayer.Players = append(scoreboardPlayer.Players, utils.Player{Name: save.Username, Score: save.Score})
		}
	}

	utils.SortPlayersByScore(scoreboardPlayer.Players)

	for i := range scoreboardPlayer.Players {
		scoreboardPlayer.Players[i].Position = i + 1
	}

	renderTemplate(w, "scoreboard/index", scoreboardPlayer)
}
