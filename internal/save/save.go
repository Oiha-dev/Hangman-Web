package save

import (
	"encoding/json"
	"fmt"
	"hangman-web/pkg/utils"
	"os"
)

func saveData(data utils.Saves) {
	file, err := os.Create("data/save.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling data:", err)
		return
	}

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Error writing data:", err)
		return
	}
}

func AddData(Username string, CurrentWord string, TestedLetters []string) {
	data := LoadData()
	data.Saves = append(data.Saves, utils.Save{
		Username:      Username,
		CurrentWord:   CurrentWord,
		TestedLetters: TestedLetters,
	})
	saveData(data)
}

func LoadData() utils.Saves {
	file, err := os.Open("data/save.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return utils.Saves{}
	}
	defer file.Close()

	var data utils.Saves
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		fmt.Println("Error decoding data:", err)
		return utils.Saves{}
	}

	return data
}
