package back

import (
	"encoding/json"
	"hangman-web/pkg/hangman-classic/structure"
	"hangman-web/pkg/utils"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func endScreen(w http.ResponseWriter, r *http.Request) {
	// Retrieve the game data cookie
	gameDataCookie, err := r.Cookie("gameData")
	if err != nil {
		http.Error(w, "Game data not found", http.StatusInternalServerError)
		return
	}

	// Unescape the game data
	gameDataValue, err := url.QueryUnescape(gameDataCookie.Value)
	if err != nil {
		http.Error(w, "Error decoding game data", http.StatusInternalServerError)
		return
	}

	// Unmarshal the game data
	var gameData structure.HangManData
	err = json.Unmarshal([]byte(gameDataValue), &gameData)
	if err != nil {
		http.Error(w, "Error parsing game data", http.StatusInternalServerError)
		return
	}

	// Retrieve the username cookie
	usernameCookie, err := r.Cookie("playerName")
	if err != nil {
		http.Error(w, "Failed to get username cookie", http.StatusBadRequest)
		return
	}

	// Save the game data into /data/save.json
	gameSave := utils.Save{
		Username:      usernameCookie.Value,
		CurrentWord:   gameData.Word,
		GoalWord:      gameData.ToFind,
		TestedLetters: gameData.GuessedLetters,
		Score:         gameData.Score,
		AttemptsLeft:  gameData.Attempts,
	}

	jsonFile, err := os.OpenFile("data/save.json", os.O_RDWR, 0644)
	if err != nil {
		http.Error(w, "Failed to open JSON file", http.StatusInternalServerError)
		return
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		http.Error(w, "Failed to read JSON file", http.StatusInternalServerError)
		return
	}

	var saves utils.Saves
	err = json.Unmarshal(byteValue, &saves)
	if err != nil {
		http.Error(w, "Failed to unmarshal JSON data", http.StatusInternalServerError)
		return
	}

	saves.Saves = append(saves.Saves, gameSave)

	updatedData, err := json.Marshal(saves)
	if err != nil {
		http.Error(w, "Failed to marshal JSON data", http.StatusInternalServerError)
		return
	}

	err = ioutil.WriteFile("data/save.json", updatedData, 0644)
	if err != nil {
		http.Error(w, "Failed to write JSON file", http.StatusInternalServerError)
		return
	}

	// Prepare data for the template
	data := map[string]interface{}{
		"Win":   gameData.IsWinned,
		"Word":  gameData.Word,
		"Score": gameData.Score,
	}

	// Parse the template
	tmpl, err := template.ParseFiles("internal/web/front/endScreen/index.gohtml")
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	// Execute the template
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}
