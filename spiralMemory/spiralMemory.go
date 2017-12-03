package spiralMemory

import (
	"fmt"
	"math"
)

// GetSteps returns the number of steps to carry the data
func GetSteps(number float64) float64 {
	c := NewCursor()
	// fmt.Printf("Number to look for: %f, cursor: %f\n", number, c.currentNumber)
	for c.currentNumber < number {
		c.next()
		// fmt.Println(c)
	}
	return math.Abs(c.currentColumn) + math.Abs(c.currentRow)
}

// NewCursor creates a new Cursor
func NewCursor() Cursor {
	c := Cursor{spiralLevel: 1, currentColumn: 0, currentRow: 0, startNumber: 1, currentNumber: 1, rowOffset: 0, columnOffset: 1, nextLevelCount: 5}
	return c
}

// Cursor is a struct
type Cursor struct {
	spiralLevel    float64
	currentColumn  float64
	currentRow     float64
	startNumber    float64
	currentNumber  float64
	columnOffset   float64
	rowOffset      float64
	nextLevelCount float64
}

func (c Cursor) String() string {
	return fmt.Sprintf("Number: %f, Current Column: %f (%f), Current Row: %f (%f), Spiral Level: %f, Next Spiral Level Count: %f", c.currentNumber, c.currentColumn, c.columnOffset, c.currentRow, c.rowOffset, c.spiralLevel, c.nextLevelCount)
}

func (c *Cursor) next() float64 {
	c.currentNumber++

	if c.nextLevelCount == 0 {
		c.spiralLevel++
		c.nextLevelCount = 5
	}

	switch {
	case c.nextLevelCount == 5:
		c.columnOffset = 1
		c.rowOffset = 0
	case c.nextLevelCount == 4:
		c.columnOffset = 0
		c.rowOffset = 1
	case c.nextLevelCount == 3:
		c.columnOffset = -1
		c.rowOffset = 0
	case c.nextLevelCount == 2:
		c.columnOffset = 0
		c.rowOffset = -1
	case c.nextLevelCount == 1:
		c.columnOffset = 1
		c.rowOffset = 0
	}

	c.currentColumn += c.columnOffset
	c.currentRow += c.rowOffset

	if c.columnOffset != 0 && math.Abs(c.currentColumn) == c.spiralLevel {
		c.nextLevelCount--
	}

	if c.rowOffset != 0 && math.Abs(c.currentRow) == c.spiralLevel {
		c.nextLevelCount--
	}

	return c.currentNumber
}
