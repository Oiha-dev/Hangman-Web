package back

import (
	"encoding/json"
	"fmt"
	"hangman-web/internal/game"
	classic_utils "hangman-web/pkg/hangman-classic/pkg/utils"
	"hangman-web/pkg/hangman-classic/structure"
	"hangman-web/pkg/utils"
	"html/template"
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
		HangmanState:   0,
		GuessedLetters: []string{},
		Score:          0,
		IsWinned:       false,
	}

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
		"split":    utils.Split,
		"contains": classic_utils.ContainsStr,
	}

	tmpl, err := template.New("index.gohtml").Funcs(funcMap).ParseFiles("internal/web/front/game/index.gohtml")
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
}

func handleGuess(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("playerName")
	if err != nil {
		http.Error(w, "Player not found", http.StatusNotFound)
		return
	}

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

	r.ParseForm()

	letterGuessed := strings.ToLower(r.URL.Query().Get("letter"))
	fullWordGuessed := strings.ToLower(r.URL.Query().Get("word"))

	var newJose structure.HangManData
	err = json.Unmarshal([]byte(gameDataValue), &newJose)
	if err != nil {
		http.Error(w, "Error decoding game data", http.StatusInternalServerError)
		return
	}

	if letterGuessed != "" {
		game.RoundLogic(&newJose, letterGuessed)
	} else if fullWordGuessed != "" {
		if strings.ToLower(newJose.Word) == fullWordGuessed {
			newJose.IsWinned = true
			newJose.ToFind = newJose.Word
		} else {
			newJose.HangmanState += 2
		}
	}

	if utils.IsFinished(newJose) {
		newJose.IsWinned = utils.IsWinned(newJose)

		gameDataJSON, err := json.Marshal(newJose)
		if err != nil {
			http.Error(w, "Error encoding game data", http.StatusInternalServerError)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:  "gameData",
			Value: url.QueryEscape(string(gameDataJSON)),
			Path:  "/",
		})

		http.Redirect(w, r, "/end", http.StatusSeeOther)
		return
	}

	gameDataJSON, err := json.Marshal(newJose)
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
		"Value":          cookie.Value,
		"AsciiArt":       utils.GetAsciiArt(newJose.HangmanState),
		"Word":           newJose.ToFind,
		"GuessedLetters": newJose.GuessedLetters,
	}

	funcMap := template.FuncMap{
		"split":    utils.Split,
		"contains": classic_utils.ContainsStr,
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

func GuessLetter(w http.ResponseWriter, r *http.Request) {
	handleGuess(w, r)
}

func FullWordGuess(w http.ResponseWriter, r *http.Request) {
	handleGuess(w, r)
}
