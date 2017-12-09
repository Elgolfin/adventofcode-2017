package streamprocess

// Score returns the score of the content string based on the rules of the Day 9
func Score(content string) int {
	score := 0
	currentGroupLevel := 0
	previousChar := ""
	isGarbage := false
	for _, r := range content {
		c := string(r)
		// If in a garbage stream, ignore the character following a !
		if isGarbage && previousChar == "!" {
			previousChar = ""
			continue
		}
		switch {
		case !isGarbage && c == "{":
			currentGroupLevel++
			score += currentGroupLevel
		case !isGarbage && c == "}":
			currentGroupLevel--
		case !isGarbage && c == "<":
			isGarbage = true
		case isGarbage && c == ">":
			isGarbage = false
		}
		previousChar = c
	}
	return score
}
