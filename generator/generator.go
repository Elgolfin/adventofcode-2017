package generator

import (
	"github.com/elgolfin/adventofcode-2017/sliceutil"
)

// Judge returns the final count after generating 40 millions pairs
func Judge(content string) int {
	mask := 65535
	genStartValues := sliceutil.Atoi(content, ",")
	genA := Generator{genStartValues[0], 16807, 2147483647, -1}
	genB := Generator{genStartValues[1], 48271, 2147483647, -1}
	finalCount := 0
	for i := 0; i < 40000000; i++ {
		valueA := genA.NextValue()
		valueB := genB.NextValue()
		if valueA&mask == valueB&mask {
			finalCount++
		}
	}
	return finalCount
}

// NextValue returns the next value calculated by the generator
func (g *Generator) NextValue() int {
	// fmt.Printf("%v\n", g)
	previousValue := g.currentValue
	if previousValue == -1 {
		previousValue = g.startValue
	}
	g.currentValue = previousValue * g.factor % g.divideBy
	return g.currentValue
}

// Generator ...
type Generator struct {
	startValue   int
	factor       int
	divideBy     int
	currentValue int
}
