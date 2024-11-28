package back

import (
	"encoding/json"
	"fmt"
	classic_utils "hangman-web/pkg/hangman-classic/pkg/utils"
	"hangman-web/pkg/utils"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func history(w http.ResponseWriter, r *http.Request) {
	usernameCookie, err := r.Cookie("playerName")
	if err != nil {
		http.Error(w, "Failed to get username cookie", http.StatusBadRequest)
		return
	}
	username := usernameCookie.Value

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

	type Save struct {
		Username    string
		Status      string
		CurrentWord string
		Score       int
		ID          int
	}

	var savesToDisplay []Save

	for i, save := range saves.Saves {
		if strings.ToLower(save.Username) != strings.ToLower(username) {
			continue
		}

		status := utils.GetSaveStatus(save)

		score := 0
		for _, letter := range save.TestedLetters {
			if classic_utils.IsLetterInWord(save.CurrentWord, letter) {
				score++
			}
		}

		savesToDisplay = append(savesToDisplay, Save{
			Username:    save.Username,
			Status:      status,
			CurrentWord: save.CurrentWord,
			Score:       score,
			ID:          i,
		})
	}

	tmpl, err := template.ParseFiles("internal/web/front/history/index.gohtml")
	if err != nil {
		fmt.Println("Failed to parse template:", err)
		return
	}

	err = tmpl.Execute(w, savesToDisplay)
	if err != nil {
		fmt.Println("Failed to execute template:", err)
	}
}
