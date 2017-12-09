package streamprocess

// Score returns the score and the number of characters within the garbage of the content string based on the rules of the Day 9
func Score(content string) (int, int) {
	score := 0
	currentGroupLevel := 0
	previousChar := ""
	isGarbage := false
	garbageSize := 0
	for _, r := range content {
		c := string(r)

		// Increment the garbage if we are in the garbage stream, we will adjust the size of the special cases afterwards
		if isGarbage {
			garbageSize++
		}

		// If in a garbage stream, ignore the character following a !
		if isGarbage && previousChar == "!" {
			previousChar = ""
			// The escape character and the escaped character do not count towards the garbage size, decrement the current size by 2
			garbageSize -= 2
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
			// The character ending the garbage does not count towards the garbage size, decrement the current size by 1
			garbageSize--
		}

		previousChar = c
	}
	return score, garbageSize
}
