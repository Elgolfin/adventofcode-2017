package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/elgolfin/adventofcode-2017/tubes"
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
	

	result, steps	:=tubes.WalkTheLine(string(content))
	fmt.Printf("Day19 --- Part One --- result is: %v\n", result)
	fmt.Printf("Day19 --- Part Two --- result is: %v\n", steps)

}
