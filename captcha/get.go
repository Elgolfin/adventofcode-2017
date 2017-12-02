package captcha

import (
	"strconv"
)

// Get solves the puzzle of the Day 1, Part One
func Get(content string) int {
	var firstChar int
	var currentChar int
	var previousChar int
	var err error
	captcha := 0
	for i, r := range content {
		// fmt.Printf("%d. %q\n", i, string(r))
		currentChar, err = strconv.Atoi(string(r))
		if err != nil {
			return -1
		}
		if i == 0 {
			firstChar = currentChar
		} else {
			if currentChar == previousChar {
				captcha += currentChar
			}
		}
		previousChar = currentChar
	}
	if currentChar == firstChar {
		captcha += currentChar
	}
	return captcha
}

// GetNew solves the puzzle of the Day 1, Part Two
func GetNew(content string) int {
	var currentChar int
	var halfwayAroundChar int
	var err error
	halfwayAroundSteps := len(content) / 2
	captcha := 0
	for i, r := range content {
		currentChar, err = strconv.Atoi(string(r))
		if err != nil {
			return -1
		}
		halfwayAroundChar, err = strconv.Atoi(string([]rune(content)[(i+halfwayAroundSteps)%len(content)]))
		// fmt.Printf("Current: %d, Next ahead of %d steps: %d\n", currentChar, halfwayAroundSteps, halfwayAroundChar)
		if err != nil {
			return -1
		}
		if currentChar == halfwayAroundChar {
			captcha += currentChar
		}
	}
	return captcha
}
