package back

import (
	"encoding/json"
	"hangman-web/pkg/hangman-classic/structure"
	"html/template"
	"net/http"
	"net/url"
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
