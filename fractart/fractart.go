package fractart

import (
	"strings"

	"github.com/elgolfin/adventofcode-2017/sliceutil"
)

// Draw returns the number of pixels on after 5 iterations
func Draw(content string, iterations int) int {
	square := InitializeSquare(".#./..#/###")
	sq := square.Split()
	rules, rulesReplacement := loadRules(content)
	// fmt.Printf("%v\n", square)
	for i := 0; i < iterations; i++ {
		// fmt.Printf("Iteration#%d\n", i+1)
		tmpSq := SquareCollection{make([][]Square, len(sq.items))}
		for j, row := range sq.items {
			tmpSq.items[j] = make([]Square, len(row))
			for k, s := range row {
				for l, rule := range rules {
					if s.DoesMatchRule(rule) {
						tmpSq.items[j][k] = Square{rulesReplacement[l].grid}
						// fmt.Printf("Found: %v\n", tmpSq.items[j][k])
						break
					}
				}
			}
		}

		// fmt.Printf("Col: %v\n", tmpSq.items)
		square = tmpSq.Merge()
		// fmt.Printf("Sq: %v\n", square)
		sq = square.Split()
		// fmt.Printf("Split: %v\n", sq)
	}
	return square.CountPixelOn()
}

func loadRules(content string) ([]Square, []Square) {
	rules := strings.Split(content, "\n")
	ruleSquares := make([]Square, len(rules))
	convertSquares := make([]Square, len(rules))
	for i, rule := range rules {
		splitRule := strings.Split(rule, " => ")
		ruleSquares[i], convertSquares[i] = InitializeSquare(splitRule[0]), InitializeSquare(splitRule[1])
	}
	return ruleSquares, convertSquares
}

// InitializeSquare returns a Square from a string
func InitializeSquare(content string) Square {
	rows := strings.Split(content, "/")
	grid := make([][]string, len(rows))
	for i, row := range rows {
		grid[i] = make([]string, len(row))
		for j, col := range row {
			grid[i][j] = string(col)
		}
	}
	return Square{grid}
}

// Flip operates a flip on a square of a size of 2 or 3
func (s *Square) Flip() {
	if len(s.grid) == 2 {
		s.grid[0][0], s.grid[0][1] = s.grid[0][1], s.grid[0][0]
		s.grid[1][0], s.grid[1][1] = s.grid[1][1], s.grid[1][0]
	}
	if len(s.grid) == 3 {
		s.grid[0][0], s.grid[0][2] = s.grid[0][2], s.grid[0][0]
		s.grid[1][0], s.grid[1][2] = s.grid[1][2], s.grid[1][0]
		s.grid[2][0], s.grid[2][2] = s.grid[2][2], s.grid[2][0]
	}
}

// Rotate operates a 90 degrees rotation on a square of a size of 2 or 3
func (s *Square) Rotate() {
	if len(s.grid) == 2 {
		s.grid[0][0], s.grid[0][1], s.grid[1][0], s.grid[1][1] = s.grid[1][0], s.grid[0][0], s.grid[1][1], s.grid[0][1]
	}
	if len(s.grid) == 3 {
		s.grid[0][0], s.grid[0][1], s.grid[0][2], s.grid[1][0], s.grid[1][2], s.grid[2][0], s.grid[2][1], s.grid[2][2] = s.grid[2][0], s.grid[1][0], s.grid[0][0], s.grid[2][1], s.grid[0][1], s.grid[2][2], s.grid[1][2], s.grid[0][2]
	}
}

// DoesMatchRule returns true if the square match the specified rule
func (s Square) DoesMatchRule(rule Square) bool {
	for i := 0; i < 4; i++ {
		s.Rotate()
		if sliceutil.Equal2DString(s.grid, rule.grid) {
			return true
		}
	}

	s.Flip()
	for i := 0; i < 4; i++ {
		s.Rotate()
		if sliceutil.Equal2DString(s.grid, rule.grid) {
			return true
		}
	}

	return false
}

// Split returns a collection of square of size 2 or 3 based on a grid of any size
func (s Square) Split() SquareCollection {
	squareSize := 2
	if len(s.grid)%2 != 0 && len(s.grid)%3 == 0 {
		squareSize = 3
	}
	size := len(s.grid) / squareSize
	sq := SquareCollection{make([][]Square, size)}
	// fmt.Printf("Splitting %v into %d squares...\n", s.grid, size)
	// i = row in the grid of pixels
	for i := 0; i < len(s.grid); i++ {
		// row = row in the collection of squares
		row := i / squareSize // (i.e. for the first row, 0/3 = 0 => new row in the square collection; 1/3 = 0; 2/3 = 0; 3/3 = 1 => new row in the square collection; etc.)
		gridRow := i % squareSize
		// fmt.Printf("Row collection %d / Row grid %d\n", row, gridRow)
		// Create a new row in the collection every squareSize of square
		if gridRow == 0 {
			sq.items[row] = make([]Square, size)
		}
		// j = column in the grid of pixels
		for j := 0; j < len(s.grid[i]); j++ {
			// col = column in the collection of squares
			col := j / squareSize // (i.e. for the first column, 0/3 = 0 => new column in the square collection; 1/3 = 0; 2/3 = 0; 3/3 = 1 => new column in the square collection; etc.)
			gridCol := j % squareSize
			// fmt.Printf("\tColumn collection %d / Column grid %d\n", col, gridCol)
			// First time in the row; first time in the column, create the Square in the current row collection (always happens when the col mod size == 0; i.e. for a size of 3 when col = 0 or col = 3, etc.)
			if gridRow == 0 && gridCol == 0 {
				sq.items[row][col] = Square{make([][]string, squareSize)}
			}
			if gridCol == 0 {
				sq.items[row][col].grid[gridRow] = make([]string, squareSize)
			}
			sq.items[row][col].grid[gridRow][gridCol] = s.grid[i][j]
			// fmt.Printf("\t%v\n", sq)
		}
	}
	return sq
}

// Merge returns a square of any size from a collection of squares
func (sq SquareCollection) Merge() Square {
	squareSize := len(sq.items[0][0].grid)
	grid := make([][]string, len(sq.items)*squareSize)
	gridRow := 0
	for v, row := range sq.items {

		if v == 0 {
			for i := 0; i < len(row)*squareSize; i++ {
				grid[gridRow+i] = make([]string, len(sq.items)*squareSize)
			}
		}

		gridCol := 0
		for _, square := range row {
			for i, currentGridRow := range square.grid {
				for j, pixel := range currentGridRow {
					grid[gridRow+i][gridCol+j] = pixel
				}
			}
			gridCol += squareSize
		}
		gridRow += squareSize
	}
	return Square{grid}
}

// CountPixelOn returns the number of pixels that are on in a square
func (s Square) CountPixelOn() int {
	pixelsOn := 0
	for _, row := range s.grid {
		for _, pixel := range row {
			if pixel == "#" {
				pixelsOn++
			}
		}
	}
	return pixelsOn
}

// Square returns ...
type Square struct {
	grid [][]string
}

// SquareCollection returns ...
type SquareCollection struct {
	items [][]Square
}
