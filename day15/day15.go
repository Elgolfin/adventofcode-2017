package main

import (
	"flag"
	"fmt"

	"github.com/elgolfin/adventofcode-2017/generator"
)

var input string

func init() {
	flag.StringVar(&input, "input", "", "The path of the input file")
}

func main() {
	flag.Parse()

	fmt.Printf("Day15 --- Part One --- result is: %v\n", generator.Judge(input))
	fmt.Printf("Day15 --- Part Two --- result is: %v\n", generator.ImpatientJudge(input))

}
