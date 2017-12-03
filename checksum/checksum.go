package checksum

import (
	"sort"
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

// GetDivisible returns the only two numbers in each row where one evenly divides the other
func GetDivisible(content string) int {
	values := strings.Split(content, "	")
	numbers := make([]int, len(values))
	for i, valueChar := range values {
		value, err := strconv.Atoi(string(valueChar))
		if err != nil {
			return -1
		}
		numbers[i] = value
	}
	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))

	for i, number := range numbers {
		currentNumberToTest := number
		numbersToTestAgain := numbers[i+1:]
		for _, numberToTestAgain := range numbersToTestAgain {
			if currentNumberToTest%numberToTestAgain == 0 {
				return currentNumberToTest / numberToTestAgain
			}
		}
	}
	return -1
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

// GenerateNew returns the checksum of a spreadsheet
func GenerateNew(spreadsheet string) int {
	lines := strings.Split(spreadsheet, "\n")
	checksum := 0
	for _, line := range lines {
		divisible := GetDivisible(line)
		checksum += divisible
	}
	return checksum
}
