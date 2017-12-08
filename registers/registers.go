package registers

import (
	"regexp"
	"strconv"
	"strings"
)

// Execute runs all the content instructions and returns the largest value in any register
func Execute(content string) (int, int) {
	instructionLines := strings.Split(content, "\n")
	registers := make(map[string]int)
	highestValueHeld := 0
	for _, instructionLine := range instructionLines {
		instruction := ParseInstructionLine(instructionLine)
		// fmt.Printf("Processing instruction %v...\n", instructionLine)
		if _, ok := registers[instruction.RegisterName]; !ok {
			registers[instruction.RegisterName] = 0
		}
		if _, ok := registers[instruction.ConditionRegisterName]; !ok {
			registers[instruction.ConditionRegisterName] = 0
		}
		currentRegisterValue := registers[instruction.RegisterName]
		conditionRegisterValue := registers[instruction.ConditionRegisterName]
		if isConditionFulfilled(instruction.ConditionOperator, conditionRegisterValue, instruction.ConditionValue) {
			registers[instruction.RegisterName] = operate(instruction.Operator, currentRegisterValue, instruction.Value)
		}
		if registers[instruction.RegisterName] > highestValueHeld {
			highestValueHeld = registers[instruction.RegisterName]
		}
		// fmt.Printf("%v\n", registers)
	}
	return getLargestValue(registers), highestValueHeld
}

func getLargestValue(registers map[string]int) int {
	largestValue := 0
	for _, v := range registers {
		if v > largestValue {
			largestValue = v
		}
	}
	return largestValue
}

func operate(operator string, registerValue int, value int) int {
	switch {
	case operator == "inc":
		registerValue += value
	case operator == "dec":
		registerValue -= value
	}
	return registerValue
}

func isConditionFulfilled(operator string, registerValue int, conditionValue int) bool {
	condition := false
	switch {
	case operator == "==":
		condition = registerValue == conditionValue
	case operator == "!=":
		condition = registerValue != conditionValue
	case operator == ">":
		condition = registerValue > conditionValue
	case operator == ">=":
		condition = registerValue >= conditionValue
	case operator == "<":
		condition = registerValue < conditionValue
	case operator == "<=":
		condition = registerValue <= conditionValue
	}
	return condition
}

// ParseInstructionLine parses an instruction line and returns corresponding properties
func ParseInstructionLine(programLine string) Instruction {
	registerName := ""
	value := 0
	operator := ""
	conditionOperator := ""
	conditionRegisterName := ""
	conditionValue := 0

	// gif inc -533 if qt <= 7
	re := regexp.MustCompile(`(?P<regiterName>[a-z]+) (?P<operator>dec|inc) (?P<value>-?\d+) if (?P<conditionRegisterName>[a-z]+) (?P<conditionOperator>==|!=|>|<|>=|<=) (?P<conditionValue>-?\d+)`)
	groupNames := re.SubexpNames()
	for _, match := range re.FindAllStringSubmatch(programLine, -1) {
		for groupIdx, groupValue := range match {
			groupName := groupNames[groupIdx]
			switch {
			case groupName == "operator":
				operator = groupValue
			case groupName == "regiterName":
				registerName = groupValue
			case groupName == "value":
				value, _ = strconv.Atoi(groupValue)
			case groupName == "conditionOperator":
				conditionOperator = groupValue
			case groupName == "conditionRegisterName":
				conditionRegisterName = groupValue
			case groupName == "conditionValue":
				conditionValue, _ = strconv.Atoi(groupValue)
			}
		}
	}
	return Instruction{operator, registerName, value, conditionOperator, conditionRegisterName, conditionValue}
}

// Instruction ...
type Instruction struct {
	Operator              string
	RegisterName          string
	Value                 int
	ConditionOperator     string
	ConditionRegisterName string
	ConditionValue        int
}
