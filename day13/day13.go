package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/elgolfin/adventofcode-2017/packetscan"
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
	fmt.Printf("Day13 --- Part One --- result is: %v\n", packetscan.GoThroughTheFirewall(string(content)))
	//fmt.Printf("Day13 --- Part Two --- result is: %v\n", n)

}
