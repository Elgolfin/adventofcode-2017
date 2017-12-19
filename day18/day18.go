package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/elgolfin/adventofcode-2017/duet"
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
	fmt.Printf("Day18 --- Part One --- result is: %v\n", duet.PlayMusicTilNonZeroRcv(string(content)))
	//fmt.Printf("Day18 --- Part Two --- result is: %v\n", permprom.BillionDance(string(content), 16))

}
