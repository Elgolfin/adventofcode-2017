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

	fmt.Printf("Day13 --- Part One --- result is: %v\n", ddefrag.CountSquares(input))
	//fmt.Printf("Day13 --- Part Two --- result is: %v\n", n)

}
