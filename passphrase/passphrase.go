package passphrase

import "strings"

// GetValidPassphrases returns the count of valid passphrases in the content string (one passphrase per line)
func GetValidPassphrases(content string) int {
	lines := strings.Split(content, "\n")
	validCount := 0
	for _, line := range lines {
		if IsValid(line) {
			validCount++
		}
	}
	return validCount
}

// IsValid returns true if the string is a valid passphrase, false otherwise
func IsValid(passphrase string) bool {
	isValid := true
	words := strings.Split(passphrase, " ")
	wordMap := make(map[string]int)
	for _, word := range words {
		if _, ok := wordMap[word]; ok {
			wordMap[word]++
			isValid = false
		} else {
			wordMap[word] = 1
		}
	}
	return isValid
}
