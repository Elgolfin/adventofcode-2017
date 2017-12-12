package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/elgolfin/adventofcode-2017/digplumb"
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
	n, progGroups := digplumb.GetAllProgGroups(string(content))
	fmt.Printf("Day12 --- Part One --- result is: %v\n", progGroups[0])
	fmt.Printf("Day12 --- Part Two --- result is: %v\n", n)

}
