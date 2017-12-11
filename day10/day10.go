package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/elgolfin/adventofcode-2017/knothash"
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

	fmt.Printf("Day10 --- Part One --- result is: %v\n", knothash.Hash(string(content), 256))
	fmt.Printf("Day10 --- Part Two --- result is: %v\n", knothash.FullHash(string(content), 256))

}
