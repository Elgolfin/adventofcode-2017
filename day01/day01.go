package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/elgolfin/adventofcode-2017/captcha"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Day01 --- Part One --- result is: %v\n", captcha.Get(string(content)))

}
