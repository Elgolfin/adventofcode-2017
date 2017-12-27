package conflagration

import (
	"regexp"
	"strconv"
	"strings"
)

// InitializeProgram returns a program struct
func InitializeProgram(id int) Program {
	registers := map[string]int{
		"a": 0,
		"b": 0,
		"c": 0,
		"d": 0,
		"e": 0,
		"f": 0,
		"g": 0,
		"h": 0,
	}
	instructions := []string{}
	cursor := 0
	return Program{id, registers, instructions, cursor}
}

// Run returns how many times the program 1 use the mulinstruction
func (p *Program) Run() int {
	res := 0
	for p.cursor < len(p.instructions) {
		cmd, arg1, arg2 := parseLine(p.instructions[p.cursor])
		p.ExecuteInstruction(cmd, arg1, arg2)
		//fmt.Printf("%d. %s %s %s\n", p.cursor, cmd, arg1, arg2)
		if cmd == "mul" {
			res++
		}
	}
	return res
}

func parseLine(line string) (string, string, string) {
	cmd := ""
	arg1 := ""
	arg2 := ""
	re := regexp.MustCompile(`(?P<cmd>[a-z]{3}) (?P<arg1>[a-z]|\d) ?(?P<arg2>(?:[a-z])|(?:-?\d+))?`)
	groupNames := re.SubexpNames()
	for _, match := range re.FindAllStringSubmatch(line, -1) {
		for groupIdx, groupValue := range match {
			groupName := groupNames[groupIdx]
			switch {
			case groupName == "cmd":
				cmd = groupValue
			case groupName == "arg1":
				arg1 = groupValue
			case groupName == "arg2":
				arg2 = groupValue
			}
		}
	}
	return cmd, arg1, arg2
}

// Load loads a set of instructions in the program struct
func (p *Program) Load(content string) {
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		p.instructions = append(p.instructions, line)
	}
}

// ExecuteInstruction executes one instruction and place the cursor to the next instruction to be executed
func (p *Program) ExecuteInstruction(cmd string, arg1 string, arg2 string) {
	switch {
	case cmd == "set":
		p.set(arg1, arg2)
	case cmd == "sub":
		p.sub(arg1, arg2)
	case cmd == "mul":
		p.mul(arg1, arg2)
	case cmd == "jnz":
		p.jump(arg1, arg2)
	}
}

func (p *Program) set(arg1 string, arg2 string) {
	if val, err := strconv.Atoi(arg2); err == nil {
		p.registers[arg1] = val
	} else {
		p.registers[arg1] = p.registers[arg2]
	}
	p.cursor++
}

func (p *Program) sub(arg1 string, arg2 string) {
	if val, err := strconv.Atoi(arg2); err == nil {
		p.registers[arg1] -= val
	} else {
		p.registers[arg1] -= p.registers[arg2]
	}
	p.cursor++
}

func (p *Program) mul(arg1 string, arg2 string) {
	if val, err := strconv.Atoi(arg2); err == nil {
		p.registers[arg1] *= val
	} else {
		p.registers[arg1] *= p.registers[arg2]
	}
	p.cursor++
}

func (p *Program) jump(arg1 string, arg2 string) {
	val1, val2 := 0, 0
	if val, err := strconv.Atoi(arg1); err == nil {
		val1 = val
	} else {
		val1 = p.registers[arg1]
	}
	if val, err := strconv.Atoi(arg2); err == nil {
		val2 = val
	} else {
		val2 = p.registers[arg2]
	}
	if val1 != 0 {
		p.cursor += val2
	} else {
		p.cursor++
	}
}

// Program ...
type Program struct {
	id           int
	registers    map[string]int
	instructions []string
	cursor       int
}
