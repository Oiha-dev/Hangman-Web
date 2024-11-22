package utils

func Contains(word string, letter string) bool {
	for _, char := range word {
		if string(char) == letter {
			return true
		}
	}
	return false
}
