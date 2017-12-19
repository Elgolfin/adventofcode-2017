package tubes

import (
	"strings"
	"unicode"
)

const pipe = '|'
const dash = '-'
const plus = '+'
const space = ' '

const left = 1
const right = 2
const top = 4
const bottom = 8

// WalkTheLine returns the path taken by the lost network packet
func WalkTheLine(content string) string {
	result := ""
	network := initializeNetwork(content)
	// fmt.Printf("%v\n", network)
	currentPosition := network.getStartingPoint()
	nextPosition, direction := network.getNextPosition(currentPosition, bottom)
	for !network.isEnd(currentPosition, nextPosition) {
		// fmt.Printf("%v %v %v", currentPosition, nextPosition, network.isEnd(currentPosition, nextPosition))
		// fmt.Printf("\t%s", string(network.grid[currentPosition.y][currentPosition.x]))
		if unicode.IsLetter(network.grid[currentPosition.y][currentPosition.x]) {
			result += string(network.grid[currentPosition.y][currentPosition.x])
		}
		currentPosition = nextPosition
		nextPosition, direction = network.getNextPosition(currentPosition, direction)
		// fmt.Printf("END: %v %v %v", currentPosition, nextPosition, network.isEnd(currentPosition, nextPosition))
	}
	result += string(network.grid[nextPosition.y][nextPosition.x])
	// fmt.Printf("\n")
	return result
}

// Looking for the starting point
func (n *network) getStartingPoint() point {
	posX := 0
	for col, char := range n.grid[0] {
		if char == pipe {
			posX = col
		}
	}
	return point{posX, 0}
}

// No error check will be done, we assume the currentPosition and the direction are good
func (n network) getNextPosition(currentPosition point, direction int) (point, int) {
	nextPosition := point{-1, -1}
	newDirection := direction
	// fmt.Printf("\n\tGet the next position of %v (going %v)\n", currentPosition, direction)
	// There is a change in the direction
	if n.grid[currentPosition.y][currentPosition.x] == plus {
		// fmt.Printf("\t\tChange of direction")
		if direction == left || direction == right {
			postop, err := currentPosition.top()
			// fmt.Printf(" to the top or bottom... ")
			if !err && n.grid[postop.y][postop.x] != space {
				nextPosition = postop
				newDirection = top
				// fmt.Printf(" top!\n")
			} else {
				nextPosition, _ = currentPosition.bottom(n.height)
				newDirection = bottom
				// fmt.Printf(" bottom!\n")
			}
		} else {
			posleft, err := currentPosition.left()
			// fmt.Printf(" to the left or right... ")
			if !err && n.grid[posleft.y][posleft.x] != space {
				nextPosition = posleft
				newDirection = left
				// fmt.Printf(" left!\n")
			} else {
				nextPosition, _ = currentPosition.right(n.width)
				newDirection = right
				// fmt.Printf(" right!\n")
			}
		}
	} else { // No change of direction
		switch direction {
		case left:
			nextPosition, _ = currentPosition.left()
		case right:
			nextPosition, _ = currentPosition.right(n.width)
		case top:
			nextPosition, _ = currentPosition.top()
		case bottom:
			nextPosition, _ = currentPosition.bottom(n.height)
		}
	}
	// fmt.Printf("\tGoing to %v (%v)\n", nextPosition, newDirection)
	return nextPosition, newDirection
}

func (n network) isEnd(currentPosition point, nextPosition point) bool {
	end := true
	posleft, posright, postop, posbottom := point{-1, -1}, point{-1, -1}, point{-1, -1}, point{-1, -1}
	var err bool
	posleft, err = nextPosition.left()
	if !err {
		if posleft != currentPosition && n.grid[posleft.y][posleft.x] != space {
			return false
		}
	}

	posright, err = nextPosition.right(n.width)
	if !err {
		if posright != currentPosition && n.grid[posright.y][posright.x] != space {
			return false
		}
	}

	postop, err = nextPosition.top()
	if !err {
		if postop != currentPosition && n.grid[postop.y][postop.x] != space {
			return false
		}
	}

	posbottom, err = nextPosition.bottom(n.height)
	if !err {
		if posbottom != currentPosition && n.grid[posbottom.y][posbottom.x] != space {
			return false
		}
	}
	return end
}

func initializeNetwork(content string) network {
	lines := strings.Split(content, "\n")
	height := len(lines)
	width := 0
	grid := make([][]rune, height)
	for row, line := range lines {
		width = len(line)
		grid[row] = make([]rune, width)
		for col, char := range line {
			grid[row][col] = char
		}
	}
	return network{grid, width, height}
}

func (p point) left() (point, bool) {
	if p.x-1 >= 0 {
		return point{p.x - 1, p.y}, false
	}
	return point{-1, -1}, true
}

func (p point) right(width int) (point, bool) {
	if p.x+1 < width {
		return point{p.x + 1, p.y}, false
	}
	return point{-1, -1}, true
}

func (p point) top() (point, bool) {
	if p.y-1 >= 0 {
		return point{p.x, p.y - 1}, false
	}
	return point{-1, -1}, true
}

func (p point) bottom(height int) (point, bool) {
	if p.y+1 < height {
		return point{p.x, p.y + 1}, false
	}
	return point{-1, -1}, true
}

type network struct {
	grid   [][]rune
	width  int
	height int
}

type point struct {
	x int
	y int
}
