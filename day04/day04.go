package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/elgolfin/adventofcode-2017/passphrase"
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

	fmt.Printf("Day04 --- Part One --- result is: %v\n", passphrase.GetValidPassphrases(string(content)))
	// fmt.Printf("Day04 --- Part Two --- result is: %v\n", passphrase.GetNew(string(content)))

}
