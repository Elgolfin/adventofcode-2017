package captcha

import "strconv"

// Get solves the puzzle of the Day 1
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
