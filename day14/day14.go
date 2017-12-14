package main

import (
	"flag"
	"fmt"

	"github.com/elgolfin/adventofcode-2017/ddefrag"
)

var input string

func init() {
	flag.StringVar(&input, "input", "", "The path of the input file")
}

func main() {
	flag.Parse()
	onesCount, grid := ddefrag.CountSquares(input)
	regions, grid := ddefrag.AreasCount(grid)
	fmt.Printf("Day14 --- Part One --- result is: %v\n", onesCount)
	fmt.Printf("Day14 --- Part Two --- result is: %v\n", regions)

}
