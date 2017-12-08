package maze

import (
	"github.com/elgolfin/adventofcode-2017/sliceutil"
)

// FindExit returns the number of steps to escape the maze of intructions
func FindExit(content string) int {
	instructions := sliceutil.Atoi(content, "\n")
	steps := 0
	offset := 0
	previousOffset := 0
	for offset < len(instructions) {
		previousOffset = offset
		offset += instructions[offset]
		instructions[previousOffset]++
		steps++
		// fmt.Printf("%d. %v\n", steps, instructions)
	}
	return steps
}

// FindExitSranger returns the number of steps to escape the maze of intructions
func FindExitSranger(content string) int {
	instructions := sliceutil.Atoi(content, "\n")
	steps := 0
	offset := 0
	previousOffset := 0
	// fmt.Printf("%d. %v\n", steps, instructions)
	for offset < len(instructions) {
		previousOffset = offset
		offset += instructions[offset]
		if instructions[previousOffset] >= 3 {
			instructions[previousOffset]--
		} else {
			instructions[previousOffset]++
		}
		steps++
		// fmt.Printf("%d. %v\n", steps, instructions)
	}
	return steps
}
