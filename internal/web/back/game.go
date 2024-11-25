package back

import (
	"hangman-web/internal/game"
	"net/http"
)

func gamePage(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("playerName")
	if err != nil {
		http.Error(w, "Player not found", http.StatusNotFound)
		return
	}

	err = game.InitPositions()
	if err != nil {
		http.Error(w, "Error initializing positions", http.StatusInternalServerError)
		return
	}

	asciiArt := game.GetNextPosition()

	data := map[string]string{
		"Value":    cookie.Value,
		"AsciiArt": asciiArt,
	}

	renderTemplate(w, "game/index", data)
}
