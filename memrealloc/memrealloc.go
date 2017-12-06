package memrealloc

import "github.com/elgolfin/adventofcode-2017/sliceutil"
import "fmt"

// MemoryBank will hold the state of 16 banks of memory
type MemoryBank struct {
	banks []int
}

// Redistribute reallocs a bank of memory until a blocks-in-banks configuration produced has already been seen before
func Redistribute(str string) (int, string, MemoryBank) {
	var memBank MemoryBank
	footPrint := make(map[string]bool)
	steps := 0
	memBank.banks = sliceutil.Atoi(str, "	")
	checksum := memBank.String()

	footPrint[checksum] = true

	for true {
		steps++
		memBank.Redistribute()
		checksum = memBank.String()
		if _, ok := footPrint[checksum]; ok {
			break
		} else {
			footPrint[checksum] = true
		}
	}

	return steps, checksum, memBank
}

// GetLoopSize reallocs a bank of memory until a blocks-in-banks configuration produced has already been seen before, and then get the size of the loop that continously produces the same comfiguration
func GetLoopSize(str string) int {
	_, checksumStep, memBank := Redistribute(str)
	steps := 0

	for true {
		steps++
		memBank.Redistribute()
		checksum := memBank.String()
		if checksumStep == checksum {
			break
		}
	}

	return steps
}

// Redistribute reallocs a bank of memory
func (mb *MemoryBank) Redistribute() {
	currentIndex := sliceutil.GetLargest(mb.banks)
	blocks := mb.banks[currentIndex]
	mb.banks[currentIndex] = 0
	for true {
		currentIndex++
		if currentIndex >= len(mb.banks) {
			currentIndex = 0
		}
		mb.banks[currentIndex]++
		blocks--
		if blocks == 0 {
			break
		}
	}
}

func (mb MemoryBank) String() string {
	resultStr := ""
	sep := ""
	for i, value := range mb.banks {
		if i > 0 {
			sep = "	"
		}
		resultStr += fmt.Sprintf("%v%v", sep, value)
	}
	return resultStr
}
