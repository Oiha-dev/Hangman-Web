package back

import (
	"encoding/json"
	"fmt"
	"hangman-web/internal/game"
	classic_utils "hangman-web/pkg/hangman-classic/pkg/utils"
	"hangman-web/pkg/hangman-classic/structure"
	"hangman-web/pkg/utils"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func gamePage(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("playerName")
	if err != nil {
		http.Error(w, "Player not found", http.StatusNotFound)
		return
	}

	randomWord := game.GetRandomWord(game.ImportWords())

	fmt.Println(randomWord)

	jose := structure.HangManData{
		Word:           randomWord,
		ToFind:         classic_utils.FirstPrintWord(randomWord),
		Attempts:       9,
		HangmanState:   1,
		GuessedLetters: []string{},
	}

	// Store the initial game state in a cookie
	gameDataJSON, err := json.Marshal(jose)
	if err != nil {
		http.Error(w, "Error encoding game data", http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "gameData",
		Value: url.QueryEscape(string(gameDataJSON)),
		Path:  "/",
	})

	data := map[string]interface{}{
		"Value":    cookie.Value,
		"AsciiArt": utils.GetAsciiArt(jose.HangmanState),
		"Word":     jose.ToFind,
	}

	funcMap := template.FuncMap{
		"split": split,
	}

	tmpl, err := template.New("index.gohtml").Funcs(funcMap).ParseFiles("internal/web/front/game/index.gohtml")
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}

func guessLetter(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("playerName")
	if err != nil {
		http.Error(w, "Player not found", http.StatusNotFound)
		return
	}

	// Retrieve the game state from the cookie
	gameDataCookie, err := r.Cookie("gameData")
	if err != nil {
		http.Error(w, "Game data not found", http.StatusInternalServerError)
		return
	}

	gameDataValue, err := url.QueryUnescape(gameDataCookie.Value)
	if err != nil {
		http.Error(w, "Error decoding game data", http.StatusInternalServerError)
		return
	}

	var jose structure.HangManData
	err = json.Unmarshal([]byte(gameDataValue), &jose)
	if err != nil {
		log.Printf("Error decoding game data: %v", err)
		log.Printf("Game data cookie value: %s", gameDataValue)
		http.Error(w, "Error decoding game data", http.StatusInternalServerError)
		return
	}

	r.ParseForm()
	guessLetter := r.FormValue("guessLetter")

	if classic_utils.IsLetterInWord(jose.ToFind, guessLetter) {
		classic_utils.UpdateWord(&jose, guessLetter)
	} else {
		jose.Attempts--
		jose.HangmanState++
	}

	// Update the game state in the cookie
	gameDataJSON, err := json.Marshal(jose)
	if err != nil {
		http.Error(w, "Error encoding game data", http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "gameData",
		Value: url.QueryEscape(string(gameDataJSON)),
		Path:  "/",
	})

	data := map[string]interface{}{
		"Value":    cookie.Value,
		"AsciiArt": utils.GetAsciiArt(jose.HangmanState),
		"Word":     jose.Word,
	}

	funcMap := template.FuncMap{
		"split": split,
	}

	tmpl, err := template.New("index.gohtml").Funcs(funcMap).ParseFiles("internal/web/front/game/index.gohtml")
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}

func split(s, sep string) []string {
	return strings.Split(s, sep)
}
