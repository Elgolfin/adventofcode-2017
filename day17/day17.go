package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/elgolfin/adventofcode-2017/spinlock"
)

var input string

func init() {
	flag.StringVar(&input, "input", "", "The path of the input file")
}

func main() {
	flag.Parse()

	n, _ := strconv.Atoi(input)
	fmt.Printf("Day17 --- Part One --- result is: %v\n", spinlock.Process(n))
	fmt.Printf("Day17 --- Part Two --- result is: %v\n", spinlock.AngryProcess(n))

}
