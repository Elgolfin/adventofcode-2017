package sliceutil

import (
	"strconv"
	"strings"
)

// Atoi initializes an array of int from a string of int separated by a character
func Atoi(content string, separator string) []int {
	lines := strings.Split(content, separator)
	sliceOfInts := make([]int, len(lines))
	for i, instruction := range lines {
		sliceOfInts[i], _ = strconv.Atoi(instruction)
	}
	return sliceOfInts
}

// Equal returns true if two slices are equal (same length, same values), false otherwise
func Equal(s1, s2 []int) bool {
	if s1 == nil && s2 == nil {
		return true
	}

	if s1 == nil || s2 == nil {
		return false
	}

	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

// GetLargest returns the index of first largest value of a slice of ints
func GetLargest(values []int) int {
	var largestValue int
	index := -1
	for i, value := range values {
		if value > largestValue {
			largestValue = value
			index = i
		}
	}
	return index
}
