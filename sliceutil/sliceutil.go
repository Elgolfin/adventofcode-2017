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

// Equal2DString returns true if two 2D slices of string are equal (same length, same values), false otherwise
func Equal2DString(s1, s2 [][]string) bool {
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
		if EqualString(s1[i], s2[i]) == false {
			return false
		}
	}

	return true

}

// Equal2DInt returns true if two 2D slices of int are equal (same length, same values), false otherwise
func Equal2DInt(s1, s2 [][]int) bool {
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
		if EqualInt(s1[i], s2[i]) == false {
			return false
		}
	}

	return true

}

// EqualString returns true if two slices of string are equal (same length, same values), false otherwise
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

// ReverseInt reverses an array of strings (i.e. [2, 7, 3, 1] becomes [1, 3, 7, 2])
func ReverseInt(input []int) []int {
	for i := 0; i < len(input)/2; i++ {
		input[i], input[len(input)-1-i] = input[len(input)-1-i], input[i]
	}
	return input
}

// InsertInt inserts an int in an array of int if it is not already present in the array
func InsertInt(ins int, arr []int) []int {
	found := false
	for _, val := range arr {
		if val == ins {
			found = true
			break
		}
	}
	if !found {
		arr = append(arr, ins)
	}
	return arr
}

// HasInt returns true if the array of ints has the value find, false otherwise
func HasInt(find int, arr []int) bool {
	for _, val := range arr {
		if val == find {
			return true
		}
	}
	return false
}

// CircularAdd returns the next index after adding n to the current index in an array of the specified length
func CircularAdd(currentIndex int, n int, length int) int {
	newIndex := 0

	if length <= 0 {
		return 0
	}

	newIndex = currentIndex + n
	if newIndex >= length {
		newIndex = newIndex % length
	}
	return newIndex
}
