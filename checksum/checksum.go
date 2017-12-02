package checksum

import (
	"strconv"
	"strings"
)

// GetSmallest returns the smallest value of a row of numbers
func GetSmallest(content string) int {
	smallestValue := 999999
	values := strings.Split(content, "	")
	for _, valueChar := range values {
		value, err := strconv.Atoi(string(valueChar))
		// fmt.Printf("Current value: %d (Smallest value: %d\n)", value, smallestValue)
		if err != nil {
			return -1
		}
		if value < smallestValue {
			smallestValue = value
		}
	}
	return smallestValue
}

// GetLargest returns the largest value of a row of numbers
func GetLargest(content string) int {
	largestValue := 0
	values := strings.Split(content, "	")
	for _, valueChar := range values {
		value, err := strconv.Atoi(string(valueChar))
		// fmt.Printf("Current value: %d (Largest value: %d)\n", value, largestValue)
		if err != nil {
			return -1
		}
		if value > largestValue {
			largestValue = value
		}
	}
	return largestValue
}

// Generate returns the checksum of a spreadsheet
func Generate(spreadsheet string) int {
	lines := strings.Split(spreadsheet, "\n")
	checksum := 0
	for _, line := range lines {
		diff := GetLargest(line) - GetSmallest(line)
		checksum += diff
	}
	return checksum
}
