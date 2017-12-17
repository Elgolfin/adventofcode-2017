package spinlock

import (
	"github.com/elgolfin/adventofcode-2017/sliceutil"
)

// Process returns
func Process(n int) int {
	circularBuffer := []int{0}
	currentPosition := 0
	// fmt.Printf("0.\t%v\n", circularBuffer)
	for i := 1; i <= 2017; i++ {
		// fmt.Printf("%d. %v", i, circularBuffer)
		nextInsertPosition := sliceutil.CircularAdd(currentPosition, n, len(circularBuffer))

		startBuffer := circularBuffer[:nextInsertPosition+1]
		endBuffer := circularBuffer[nextInsertPosition+1:]

		if nextInsertPosition == 0 {
			startBuffer = circularBuffer[:1]
			endBuffer = circularBuffer[1:]
		}
		if nextInsertPosition+1 > len(circularBuffer) {
			startBuffer = circularBuffer[:]
			endBuffer = circularBuffer[len(circularBuffer):]
		}

		tmpStartBuffer := make([]int, len(startBuffer))
		tmpEndBuffer := make([]int, len(endBuffer))

		copy(tmpStartBuffer, startBuffer)
		copy(tmpEndBuffer, endBuffer)

		tmpStartBuffer = append(tmpStartBuffer, i)
		circularBuffer = append(tmpStartBuffer, tmpEndBuffer...)
		currentPosition = len(tmpStartBuffer) - 1
	}
	afterPos := currentPosition + 1
	if currentPosition >= len(circularBuffer) {
		afterPos = 0
	}
	return circularBuffer[afterPos]
}
