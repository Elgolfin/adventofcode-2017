package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/elgolfin/adventofcode-2017/registers"
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

	out1, out2 := registers.Execute(string(content))
	fmt.Printf("Day08 --- Part One --- result is: %v\n", out1)
	fmt.Printf("Day08 --- Part Two --- result is: %v\n", out2)

}
