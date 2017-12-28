package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/elgolfin/adventofcode-2017/turing"
)

var input string

func init() {
	flag.StringVar(&input, "input", "", "The path of the input file")
}

func main() {
	flag.Parse()

	steps, _ := strconv.Atoi(input)

	fmt.Printf("Day25 --- Part One --- result is: %v\n", turing.Run(steps))
}
