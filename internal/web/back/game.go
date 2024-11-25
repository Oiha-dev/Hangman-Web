package back

import (
	"hangman-web/pkg/utils"
	"net/http"
)

func gamePage(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("playerName")
	if err != nil {
		http.Error(w, "Player not found", http.StatusNotFound)
		return
	}

	asciiArt := utils.GetAsciiArt(9)

	data := map[string]string{
		"Value":    cookie.Value,
		"AsciiArt": asciiArt,
	}

	renderTemplate(w, "game/index", data)
}
