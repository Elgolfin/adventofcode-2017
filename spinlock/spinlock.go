package spinlock

import (
	"github.com/elgolfin/adventofcode-2017/sliceutil"
)

// Process returns the value after 2017
func Process(n int) int {
	circularBuffer := []int{0}
	currentPosition := 0
	for i := 1; i <= 2017; i++ {
		nextInsertPosition := sliceutil.CircularAdd(currentPosition, n, len(circularBuffer)) + 1

		startBuffer := circularBuffer[:nextInsertPosition]
		endBuffer := circularBuffer[nextInsertPosition:]

		if nextInsertPosition > len(circularBuffer) {
			startBuffer = circularBuffer[:]
			endBuffer = []int{}
		}

		tmpStartBuffer := make([]int, len(startBuffer))
		tmpEndBuffer := make([]int, len(endBuffer))

		copy(tmpStartBuffer, startBuffer)
		copy(tmpEndBuffer, endBuffer)

		tmpStartBuffer = append(tmpStartBuffer, i)
		circularBuffer = append(tmpStartBuffer, tmpEndBuffer...)
		currentPosition = nextInsertPosition
	}
	afterPos := currentPosition + 1
	if currentPosition >= len(circularBuffer) {
		afterPos = 0
	}
	return circularBuffer[afterPos]
}

// AngryProcess returns the value after 0 the moment 50,000,000 is inserted
func AngryProcess(n int) int {
	currentPosition := 0
	valueAfter0 := 0
	for i := 1; i <= 50000000; i++ {
		nextInsertPosition := sliceutil.CircularAdd(currentPosition, n, i) + 1

		// The value 0 will always be at index 0, so the value after 0 will always be inserted at index 1
		if nextInsertPosition == 1 {
			valueAfter0 = i
		}

		currentPosition = nextInsertPosition
	}

	return valueAfter0
}
