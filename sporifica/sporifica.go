package sporifica

import (
	"fmt"
	"strings"
)

const infected = "#"
const clean = "."
const flagged = "F"
const weakened = "W"

const up = 1
const down = 2
const left = 4
const right = 8

// Burst returns how many nodes have become infected after n bursts of activity
func Burst(content string, n int) int {
	grid := loadGrid(content)
	infectedNodesCount := 0
	virusCarrier := virusCarrier{0, 0, up}

	for i := 1; i <= n; i++ {
		virusPos := virusCarrier.getPosition()
		currentCell := grid[virusPos]
		if currentCell == infected {
			virusCarrier.turnRight()
			grid[virusPos] = clean
		} else {
			virusCarrier.turnLeft()
			grid[virusPos] = infected
			infectedNodesCount++
		}
		virusCarrier.move()
	}

	// fmt.Printf("\n")
	// for y := 4; y >= -4; y-- {
	// 	for x := -4; x <= 4; x++ {
	// 		pos := fmt.Sprintf("%d,%d", x, y)
	// 		if grid[pos] == "" {
	// 			grid[pos] = "."
	// 		}
	// 		fmt.Printf("%s ", grid[pos])
	// 	}
	// 	fmt.Printf("\n")
	// }

	return infectedNodesCount
}

// BurstV2 returns how many nodes have become infected after n bursts of activity
func BurstV2(content string, n int) int {
	grid := loadGrid(content)
	infectedNodesCount := 0
	virusCarrier := virusCarrier{0, 0, up}

	for i := 1; i <= n; i++ {
		virusPos := virusCarrier.getPosition()
		currentCell := grid[virusPos]
		switch currentCell {
		case infected:
			virusCarrier.turnRight()
		case clean:
			virusCarrier.turnLeft()
		case flagged:
			virusCarrier.reverse()
		case weakened:
			// Do nothing
		case "":
			virusCarrier.turnLeft()
		}
		grid[virusPos] = changeNode(currentCell)
		// fmt.Printf("Node %s was %s, now is %s\n", virusPos, currentCell, grid[virusPos])
		if grid[virusPos] == infected {
			infectedNodesCount++
		}
		virusCarrier.move()
	}

	// fmt.Printf("\n")
	// for y := 4; y >= -4; y-- {
	// 	for x := -4; x <= 4; x++ {
	// 		pos := fmt.Sprintf("%d,%d", x, y)
	// 		if grid[pos] == "" {
	// 			grid[pos] = "."
	// 		}
	// 		fmt.Printf("%s ", grid[pos])
	// 	}
	// 	fmt.Printf("\n")
	// }

	return infectedNodesCount
}

// - Clean nodes become weakened.
// - Weakened nodes become infected.
// - Infected nodes become flagged.
// - Flagged nodes become clean.
func changeNode(current string) string {
	new := ""
	switch current {
	case clean:
		new = weakened
	case weakened:
		new = infected
	case infected:
		new = flagged
	case flagged:
		new = clean
	case "":
		new = weakened
	}
	return new
}

func loadGrid(content string) map[string]string {
	grid := make(map[string]string)
	rows := strings.Split(content, "\n")
	offsetY := (len(rows) - 1) / 2
	for _, row := range rows {
		offsetX := (len(row) - 1) / -2 // -2 because we start in the neagtive x-axis
		for _, cell := range row {
			pos := fmt.Sprintf("%d,%d", offsetX, offsetY)
			grid[pos] = string(cell)
			offsetX++
		}
		offsetY--
	}
	return grid
}

func (p virusCarrier) getPosition() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}

func (p *virusCarrier) turnRight() {
	switch p.dir {
	case left:
		p.dir = up
	case right:
		p.dir = down
	case up:
		p.dir = right
	case down:
		p.dir = left
	}
}

func (p *virusCarrier) turnLeft() {
	switch p.dir {
	case left:
		p.dir = down
	case right:
		p.dir = up
	case up:
		p.dir = left
	case down:
		p.dir = right
	}
}

func (p *virusCarrier) reverse() {
	switch p.dir {
	case left:
		p.dir = right
	case right:
		p.dir = left
	case up:
		p.dir = down
	case down:
		p.dir = up
	}
}

func (p *virusCarrier) move() {
	switch p.dir {
	case left:
		p.x--
	case right:
		p.x++
	case up:
		p.y++
	case down:
		p.y--
	}
}

type virusCarrier struct {
	x   int
	y   int
	dir int
}
