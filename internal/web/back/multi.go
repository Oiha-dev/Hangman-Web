package back

import (
	"encoding/json"
	"hangman-web/internal/game"
	classic_utils "hangman-web/pkg/hangman-classic/pkg/utils"
	"hangman-web/pkg/utils"
	"net/http"
	"net/url"
	"time"
)

func multi(w http.ResponseWriter, r *http.Request) {
	createCookie(w, r)
	renderTemplate(w, "game/index", nil)
}

func createCookie(w http.ResponseWriter, r *http.Request) {
	/*
		This function is used to create the cookie for the game data
	*/
	name, err := r.Cookie("playerName")
	if err != nil {
		http.Error(w, "Player name cookie not found", http.StatusBadRequest)
		return
	}
	listNames = append(listNames, name.Value)

	time.Sleep(1 * time.Second)

	easyWords, mediumWords, hardWords := game.ImportWords()
	randomWord := game.GetRandomWord(easyWords, mediumWords, hardWords, "medium")

	hangmanData := utils.HangManDataMulti{
		Word:           randomWord,
		ToFind:         classic_utils.FirstPrintWord(randomWord),
		Attempts:       9,
		HangmanState:   0,
		GuessedLetters: []string{},
		GuessedWords:   []string{},
		Score:          0,
		IsWinned:       false,
		Player1:        utils.Player{Name: listNames[0], Score: 0, Position: 0},
		Player2:        utils.Player{Name: listNames[1], Score: 0, Position: 0},
		CurrentPlayer:  1,
	}

	jsonData, err := json.Marshal(hangmanData)
	if err != nil {
		http.Error(w, "Error creating JSON data", http.StatusInternalServerError)
		userCountMutex.Unlock()
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "hangmanMulti",
		Value:    url.QueryEscape(string(jsonData)),
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
	})
}
