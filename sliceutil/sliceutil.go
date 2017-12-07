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

// EqualInt returns true if two slices of int are equal (same length, same values), false otherwise
func EqualInt(s1, s2 []int) bool {
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

// EqualString returns true if two slices of int are equal (same length, same values), false otherwise
func EqualString(s1, s2 []string) bool {
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

// ExtendString extends a slice by adding one item
func ExtendString(slice []string, element string) []string {
	n := len(slice)
	if n == cap(slice) {
		// Slice is full; must grow.
		// We double its size and add 1, so if the size is zero we still grow.
		newSlice := make([]string, len(slice), 2*len(slice)+1)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}
