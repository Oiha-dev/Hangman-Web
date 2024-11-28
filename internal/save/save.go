package save

import (
	"encoding/json"
	"fmt"
	"hangman-web/pkg/utils"
	"os"
)

func saveData(data utils.Saves) {
	/*
		This function is used to save the data in the file data/save.json
		params: the data to save
	*/
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
