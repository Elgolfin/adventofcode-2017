package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/elgolfin/adventofcode-2017/conflagration"
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
	program := conflagration.InitializeProgram(0)
	program.Load(string(content))
	fmt.Printf("Day23 --- Part One --- result is: %v\n", program.Run())
	//fmt.Printf("Day18 --- Part Two --- result is: %v\n", conflagration.Run(string(content)))

}
