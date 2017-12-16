package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/elgolfin/adventofcode-2017/permprom"
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
	fmt.Printf("Day16 --- Part One --- result is: %v\n", permprom.Dance(string(content), 16))
	// fmt.Printf("Day16 --- Part Two --- result is: %v\n", packetscan.GoSafelyThroughTheFirewall(string(content)))

}
