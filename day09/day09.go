package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/elgolfin/adventofcode-2017/streamprocess"
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

	score, garbageSize := streamprocess.Score(string(content))
	fmt.Printf("Day09 --- Part One --- result is: %v\n", score)
	fmt.Printf("Day09 --- Part Two --- result is: %v\n", garbageSize)

}
