package passphrase

import (
	"sort"
	"strings"
)

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

// GetValidPassphrasesNew returns the count of valid passphrases in the content string (one passphrase per line)
func GetValidPassphrasesNew(content string) int {
	lines := strings.Split(content, "\n")
	validCount := 0
	for _, line := range lines {
		if IsValidNew(line) {
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

// IsValidNew returns true if the string is a valid passphrase, false otherwise
func IsValidNew(passphrase string) bool {
	words := strings.Split(passphrase, " ")
	wordMap := make(map[string]bool)
	for _, word := range words {
		if _, ok := wordMap[word]; !ok {
			for key := range wordMap {
				if IsAnagram(word, key) {
					return false
				}
			}
			wordMap[word] = true
		} else { // the word already exists in the map, so it is an anagram
			return false
		}
	}
	return true
}

// IsAnagram return true if s1 is an anagram of s2
func IsAnagram(s1 string, s2 string) bool {
	// If the two strings have the same number of characters and there are identical after sorting them, this is an anagram
	return len(s1) == len(s2) && sortString(s1) == sortString(s2)
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
