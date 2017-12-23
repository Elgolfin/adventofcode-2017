package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/elgolfin/adventofcode-2017/partswarm"
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

	fmt.Printf("Day20 --- Part One --- result is: %v\n", partswarm.GetClosestParticleToOrigin(string(content)))
	// fmt.Printf("Day20 --- Part Two --- result is: %v\n", fractart.Draw(string(content), 18))

}
