package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/elgolfin/adventofcode-2017/electromoat"
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

	heaviestBridge, longestHeaviestBridge := electromoat.FindHeaviestBridges(string(content))

	fmt.Printf("Day24 --- Part One --- result is: %v\n", heaviestBridge)
	fmt.Printf("Day24 --- Part Two --- result is: %v\n", longestHeaviestBridge)
}
