package knothash

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/elgolfin/adventofcode-2017/sliceutil"
)

// Hash computes the knot hash
func Hash(content string, size int) int {
	sequence := sliceutil.Atoi(content, ",")
	circleKnots := GetStringCircle(size)
	for _, n := range sequence {
		circleKnots.ComputeKnot(n)
	}
	return circleKnots.GetHash()
}

// FullHash fully computes the knot hash
func FullHash(content string, size int) string {
	var buffer bytes.Buffer
	asciiSequence := []byte(content)
	buffer.WriteString(convert(asciiSequence))
	if buffer.Len() > 0 {
		buffer.WriteString(",17,31,73,47,23")
	} else {
		buffer.WriteString("17,31,73,47,23")
	}
	sequence := sliceutil.Atoi(buffer.String(), ",")
	circleKnots := GetStringCircle(size)
	for i := 0; i < 64; i++ {
		for _, n := range sequence {
			circleKnots.ComputeKnot(n)
		}
	}
	var knotHash = ""
	for i := 0; i < 256; i++ {
		var hexaBlock = circleKnots.list[i]
		for j := i + 1; j < i+16; j++ {
			hexaBlock ^= circleKnots.list[j]
		}
		knotHash = fmt.Sprintf("%s%02x", knotHash, hexaBlock)
		i += 15
	}
	return knotHash
}

func convert(b []byte) string {
	s := make([]string, len(b))
	for i := range b {
		s[i] = strconv.Itoa(int(b[i]))
	}
	return strings.Join(s, ",")
}

// GetStringCircle returns a StringCircle of a size of size strings properly initialized
func GetStringCircle(size int) StringCircle {
	list := make([]int, size)
	for i := range list {
		list[i] = i
	}
	return StringCircle{list, 0, 0}
}

// StringCircle ...
type StringCircle struct {
	list            []int
	currentPosition int
	skipSize        int
}

// GetHash returns the hash of the StringCircle (first value of the list times second value of the list)
func (sc StringCircle) GetHash() int {
	return sc.list[0] * sc.list[1]
}

// ComputeKnot selects a span of string, brings the ends together, and gives the span a half-twist to reverse the order of the marks within it
func (sc *StringCircle) ComputeKnot(n int) {
	// fmt.Printf("Compute knot %v for %v\n", sc, n)
	i := 0
	tmpValues := make([]int, n)  // Keep track of which values of the indexes we are working with
	tmpIndexes := make([]int, n) // Keep track of which indexes of the list we are working with (they will be used to reverse them and get the values form the tmpValues)
	currentIndex := sc.currentPosition
	for i < n {
		tmpIndexes[i] = currentIndex
		tmpValues[i] = sc.list[currentIndex]
		currentIndex++
		if currentIndex >= len(sc.list) {
			currentIndex = 0
		}
		i++
	}
	// fmt.Printf("Indexes: %v\n", tmpIndexes)
	// fmt.Printf("Values: %v\n", tmpValues)

	// Reverse the values
	tmpValues = sliceutil.ReverseInt(tmpValues)
	// fmt.Printf("Reversed Values: %v\n", tmpValues)

	// Put the back the reversed values in the list
	for j, index := range tmpIndexes {
		sc.list[index] = tmpValues[j]
	}

	// Update the current position of the list
	sc.currentPosition = sc.currentPosition + (n+sc.skipSize)%256
	if sc.currentPosition >= len(sc.list) {
		sc.currentPosition -= len(sc.list)
	}

	// Update the skip size of the list
	sc.skipSize++
}
