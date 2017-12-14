package ddefrag

import (
	"fmt"
	"math/bits"
	"strconv"

	"github.com/elgolfin/adventofcode-2017/knothash"
)

// CountSquares returns the number of squares set to 1 in a grid
func CountSquares(content string) (int, [][]int) {
	squares := 0
	grid := [][]int{}
	for r := 0; r < 128; r++ {
		grid = append(grid, []int{})
		hashInput := fmt.Sprintf("%s-%d", content, r)
		hash := knothash.FullHash(hashInput, 256)
		// fmt.Printf("%s: %d\n", hash, OnesCount(hash))
		count, row := OnesCount(hash)
		for _, c := range row {
			num, _ := strconv.Atoi(string(c))
			grid[r] = append(grid[r], num)
			// fmt.Printf("%d ", num)
		}
		// fmt.Printf("\n")
		squares += count
	}
	return squares, grid
}

// OnesCount returns the number of ones in an hexa string
func OnesCount(hexa string) (int, string) {
	count := 0
	binaryRepresentation := ""
	for _, c := range hexa {
		n, _ := strconv.ParseUint(string(c), 16, 32)
		mask := uint32(8)
		for i := 0; i < 4; i++ {
			bit := 0
			if mask&uint32(n) > 0 {
				bit = 1
			}
			mask = mask >> 1
			binaryRepresentation = fmt.Sprintf("%s%d", binaryRepresentation, bit)
		}
		count += bits.OnesCount32(uint32(n))
	}
	return count, binaryRepresentation
}

// AreasCount returns the number of regions in a grid
func AreasCount(grid [][]int) (int, [][]int) {
	region := 1
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == 1 {
				region++
				grid = FloodFill(x, y, grid, region)
			}
		}
	}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			// fmt.Printf("%d ", grid[y][x])
		}
		// fmt.Printf("\n")
	}

	return region - 1, grid
}

// FloodFill returns a grid and replace all the adjacent squares whose value is 1 of the startNode by the number provided in the replacementSquare (0,0 is the top right square)
func FloodFill(startNodeX int, startNodeY int, grid [][]int, replacementSquare int) [][]int {
	// fmt.Printf("Processing node %d,%d...\n", startNodeX, startNodeY)
	if grid[startNodeY][startNodeX] == replacementSquare {
		return grid
	}
	if grid[startNodeY][startNodeX] == 0 {
		return grid
	}
	if grid[startNodeY][startNodeX] != 1 {
		return grid
	}
	grid[startNodeY][startNodeX] = replacementSquare
	if startNodeX-1 >= 0 {
		// fmt.Printf("\t1. ")
		grid = FloodFill(startNodeX-1, startNodeY, grid, replacementSquare)
	}
	if startNodeX+1 < len(grid) {
		// fmt.Printf("\t2. ")
		grid = FloodFill(startNodeX+1, startNodeY, grid, replacementSquare)
	}
	if startNodeY-1 >= 0 {
		// fmt.Printf("\t3. ")
		grid = FloodFill(startNodeX, startNodeY-1, grid, replacementSquare)
	}
	if startNodeY+1 < len(grid) {
		// fmt.Printf("\t4. ")
		grid = FloodFill(startNodeX, startNodeY+1, grid, replacementSquare)
	}
	return grid
}
