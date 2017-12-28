package turing

import (
	"fmt"
)

const left = -1
const right = 1

// Run returns the cheksum
func Run(steps int) int {

	stepsDict := map[string]State{
		"A0": State{1, 1, "B"},
		"A1": State{0, -1, "B"},
		"B0": State{0, 1, "C"},
		"B1": State{1, -1, "B"},
		"C0": State{1, 1, "D"},
		"C1": State{0, -1, "A"},
		"D0": State{1, -1, "E"},
		"D1": State{1, -1, "F"},
		"E0": State{1, -1, "A"},
		"E1": State{0, -1, "D"},
		"F0": State{1, 1, "A"},
		"F1": State{1, -1, "E"},
	}

	slot := make(map[int]int)
	currentSlotIndex := 0
	currentStep := "A"

	for step := 0; step < steps; step++ {
		if _, ok := slot[currentSlotIndex]; !ok {
			slot[currentSlotIndex] = 0
		}
		state := stepsDict[fmt.Sprintf("%s%d", currentStep, slot[currentSlotIndex])]
		slot[currentSlotIndex] = state.writeValue
		currentSlotIndex += state.moveTo
		currentStep = state.nextStep
	}

	checksum := 0
	for _, v := range slot {
		checksum += v
	}

	return checksum
}

// State ...
type State struct {
	writeValue int
	moveTo     int
	nextStep   string
}
