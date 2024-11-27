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

	difficultyCookie, err := r.Cookie("difficulty")
	if err != nil {
		http.Error(w, "Difficulty not found", http.StatusNotFound)
		return
	}

	easyWords, mediumWords, hardWords := game.ImportWords()
	randomWord := game.GetRandomWord(easyWords, mediumWords, hardWords, difficultyCookie.Value)

	fmt.Println(cookie.Value, ":", randomWord)

	jose := structure.HangManData{
		Word:           randomWord,
		ToFind:         classic_utils.FirstPrintWord(randomWord),
		Attempts:       9,
		HangmanState:   0,
		GuessedLetters: []string{},
		GuessedWords:   []string{},
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

	letterGuessed := strings.ToLower(r.FormValue("letter"))
	if len(letterGuessed) > 1 {
		letterGuessed = letterGuessed[:1]
	}
	fullWordGuessed := strings.ToLower(r.FormValue("fullWord"))

	var newJose structure.HangManData
	err = json.Unmarshal([]byte(gameDataValue), &newJose)
	if err != nil {
		http.Error(w, "Error decoding game data", http.StatusInternalServerError)
		return
	}

	if !classic_utils.ContainsStr(newJose.GuessedLetters, strings.ToUpper(letterGuessed)) {
		if letterGuessed != "" {
			game.RoundLogic(&newJose, letterGuessed)
		} else if fullWordGuessed != "" {
			if !classic_utils.ContainsStr(newJose.GuessedWords, strings.ToUpper(fullWordGuessed)) {
				newJose.GuessedWords = append(newJose.GuessedWords, strings.ToUpper(fullWordGuessed))
				if strings.ToLower(newJose.Word) == fullWordGuessed {
					newJose.IsWinned = true
					newJose.ToFind = newJose.Word
					newJose.Score += 2
				} else {
					newJose.HangmanState += 2
					newJose.Attempts -= 2
				}
			}
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
		"GuessedWords":   newJose.GuessedWords,
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
