package duet

import (
	"regexp"
	"strconv"
	"strings"
)

// PlayMusicTilNonZeroRcv loads a program and plays its music until the first time a rcv instruction is executed with a non-zero value
func PlayMusicTilNonZeroRcv(content string) int {
	prog := Program{make(map[string]int), []string{}, 0, 0, 0}
	prog.Load(content)
	return prog.RunTilNonZeroRcv()
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
	// fmt.Printf("Execute %s %s %s\n", cmd, arg1, arg2)
	switch {
	case cmd == "snd":
		if val, err := strconv.Atoi(arg1); err == nil {
			p.lastSoundPlayedFrequency = val
		} else {
			p.lastSoundPlayedFrequency = p.registers[arg1]
		}
		p.cursor++
	case cmd == "set":
		if val, err := strconv.Atoi(arg2); err == nil {
			p.registers[arg1] = val
		} else {
			p.registers[arg1] = p.registers[arg2]
		}
		p.cursor++
	case cmd == "add":
		if val, err := strconv.Atoi(arg2); err == nil {
			p.registers[arg1] += val
		} else {
			p.registers[arg1] += p.registers[arg2]
		}
		p.cursor++
	case cmd == "mul":
		if val, err := strconv.Atoi(arg2); err == nil {
			p.registers[arg1] *= val
		} else {
			p.registers[arg1] *= p.registers[arg2]
		}
		p.cursor++
	case cmd == "mod":
		if val, err := strconv.Atoi(arg2); err == nil {
			p.registers[arg1] = p.registers[arg1] % val
		} else {
			p.registers[arg1] = p.registers[arg1] % p.registers[arg2]
		}
		p.cursor++
	case cmd == "rcv":
		tmpVal := 0
		if val, err := strconv.Atoi(arg1); err == nil {
			tmpVal = val
		} else {
			tmpVal = p.registers[arg1]
		}
		if tmpVal != 0 {
			p.recoveredFrequency = p.lastSoundPlayedFrequency
		}
		p.cursor++
	case cmd == "jgz":
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
		if val1 > 0 {
			p.cursor += val2
		} else {
			p.cursor++
		}
	}
	// fmt.Printf("%v %v %v %v\n", p.cursor, p.registers, p.lastSoundPlayedFrequency, p.recoveredFrequency)
}

// RunTilNonZeroRcv executes the instructions loaded in the program
func (p *Program) RunTilNonZeroRcv() int {
	res := 0
	for p.cursor < len(p.instructions) {
		cmd, arg1, arg2 := parseLine(p.instructions[p.cursor])
		p.ExecuteInstruction(cmd, arg1, arg2)
		if cmd == "rcv" && p.registers[arg1] != 0 {
			res = p.recoveredFrequency
			break
		}
	}
	return res
}

// Program ...
type Program struct {
	registers                map[string]int
	instructions             []string
	cursor                   int
	recoveredFrequency       int
	lastSoundPlayedFrequency int
}
