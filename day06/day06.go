package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/elgolfin/adventofcode-2017/memrealloc"
)

var input string

func init() {
	flag.StringVar(&input, "input", "", "The path of the input file")
}

func main() {
	flag.Parse()

	content, err := ioutil.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}

	steps, _, _ := memrealloc.Redistribute(string(content))
	fmt.Printf("Day06 --- Part One --- result is: %v\n", steps)
	fmt.Printf("Day06 --- Part Two --- result is: %v\n", memrealloc.GetLoopSize(string(content)))

}
