package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"

	"github.com/elgolfin/adventofcode-2017/spiralMemory"
)

var input string

func init() {
	flag.StringVar(&input, "input", "", "The input of the puzzle")
}

func main() {
	flag.Parse()

	inputFloat, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Day03 --- Part One --- result is: %v\n", spiralMemory.GetSteps(inputFloat))
	fmt.Printf("Day03 --- Part Two --- result is: %v\n", spiralMemory.GetFirstValueLargerThan(inputFloat))

}
