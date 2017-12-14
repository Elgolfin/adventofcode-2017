package ddefrag

import (
	"fmt"
	"math/bits"
	"strconv"

	"github.com/elgolfin/adventofcode-2017/knothash"
)

// CountSquares returns -1
func CountSquares(content string) int {
	squares := 0
	for r := 0; r < 128; r++ {
		hashInput := fmt.Sprintf("%s-%d", content, r)
		hash := knothash.FullHash(hashInput, 256)
		// fmt.Printf("%s: %d\n", hash, OnesCount(hash))
		squares += OnesCount(hash)
	}
	return squares
}

// OnesCount returns the number of ones in an hexa string
func OnesCount(hexa string) int {
	count := 0
	for _, c := range hexa {
		n, _ := strconv.ParseUint(string(c), 16, 32)
		count += bits.OnesCount32(uint32(n))
	}
	return count
}
