package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/elgolfin/adventofcode-2017/recurscircus"
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

	fmt.Printf("Day07 --- Part One --- result is: %v\n", recurscircus.GetRootProgramName(string(content)))
	neededWeightToBalance, _ := recurscircus.BalanceWeight(string(content))
	fmt.Printf("Day07 --- Part Two --- result is: %v\n", neededWeightToBalance)

}
