package recurscircus

import (
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/elgolfin/adventofcode-2017/sliceutil"
)

// GetRootProgramName returns the root program
func GetRootProgramName(content string) string {
	rootProgram := ""

	programsMap := constructMap(content)

	// Clean-up the map by removing each program whose all the above programs are not in the map
	// Continue until there is just one remaining program (the root one)
	for len(programsMap) > 1 {
		programsToBeRemoved := make([]string, 0)
		for progName, program := range programsMap {
			count := len(program.abovePrograms)
			for _, aboveProgram := range program.abovePrograms {
				if _, ok := programsMap[aboveProgram]; !ok {
					count--
				} else {
					rootProgram = progName
					break
				}
			}
			if count == 0 {
				programsToBeRemoved = sliceutil.ExtendString(programsToBeRemoved, progName)
			}
		}
		for _, progToBeRemoved := range programsToBeRemoved {
			delete(programsMap, progToBeRemoved)
		}
	}

	return programsMap[rootProgram].Name
}

// GetRootProgram returns the root program
func GetRootProgram(content string) Program {
	rootProgramName := GetRootProgramName(content)
	programsMap := constructMap(content)
	return *programsMap[rootProgramName]
}

// BalanceWeight returns the weight of the program should be to balance the whole tower
func BalanceWeight(content string) (int, *Program) {
	rootProg := GetRootProgram(content)
	unbalancedProg := &rootProg
	neededWeightToBalance := -1
	balancedWeight := 0
	unbalancedWeight := 0
	for {
		// fmt.Printf("%s has %d children:", unbalancedProg.Name, len(unbalancedProg.ptrAbovePrograms))
		weights := make(map[int][]*Program)
		weight := 0
		for _, aboveProgram := range unbalancedProg.ptrAbovePrograms {
			// fmt.Printf("Getting the weight of the program %s...", aboveProgramName)
			weight = aboveProgram.GetWeight()
			weights[weight] = append(weights[weight], aboveProgram)
		}

		// All the above programs share the same weight, the parent program is the one that needs to balance its weight (see previous loop iteration)
		if len(weights) == 1 {
			// fmt.Printf(" they are all balanced (%d)\n", weight)
			break
		}

		for k, v := range weights {
			if len(v) == 1 {
				unbalancedWeight = k
				unbalancedProg = v[0]
			} else {
				balancedWeight = k
			}
		}
		// fmt.Printf(" one (%s/%d) is unbalanced, the others are balanced (%d)\n", unbalancedProg.Name, unbalancedWeight, balancedWeight)
		neededWeightToBalance = int(math.Abs(float64(unbalancedWeight) - float64(balancedWeight)))
	}

	if unbalancedProg.Weight > balancedWeight {
		neededWeightToBalance = unbalancedProg.Weight + neededWeightToBalance
	} else {
		neededWeightToBalance = unbalancedProg.Weight - neededWeightToBalance
	}

	// fmt.Printf("Unbalanced program: %s\n", unbalancedProg.Name)
	return neededWeightToBalance, unbalancedProg
}

// GetWeight returns the weight of the sub-tower's program
func (program Program) GetWeight() int {
	weight := program.Weight
	// fmt.Printf("Getting the weight of the program %s (%d)...\n", program.Name, program.weight)
	if len(program.ptrAbovePrograms) > 0 {
		for _, aboveProgram := range program.ptrAbovePrograms {
			// fmt.Printf("\tGetting the weight of the above program %s...\n", aboveProgram.Name)
			weight += aboveProgram.GetWeight()
			// fmt.Printf("... %d. Done.\n", weight)
		}
	}
	return weight
}

func constructMap(content string) map[string]*Program {
	programsLine := strings.Split(content, "\n")
	programsMap := make(map[string]*Program)
	for _, programLine := range programsLine {
		newProg := ParseProgramLine(programLine)
		program := &newProg
		// fmt.Printf("%d. %s <=> %v\n", i, programLine, program)
		if _, ok := programsMap[program.Name]; !ok {
			programsMap[program.Name] = program
			// fmt.Printf("\t%s does not exist. Creating it (%p)\n", program.Name, &program)
		} else {
			program = programsMap[program.Name]
			program.Weight = newProg.Weight
			program.abovePrograms = newProg.abovePrograms
			// fmt.Printf("\t%s already exists %v (%p)\n", program.Name, program, program)
		}

		for _, aboveProgName := range program.abovePrograms {
			if val, ok := programsMap[aboveProgName]; !ok {
				aboveProgram := Program{aboveProgName, -1, nil, nil}
				program.ptrAbovePrograms = append(program.ptrAbovePrograms, &aboveProgram)
				// fmt.Printf("\t\t%s does not exist. Creating it (%p)\n", aboveProgName, &aboveProgram)
				programsMap[aboveProgName] = &aboveProgram
			} else {
				// fmt.Printf("\t\t%s already exists %v (%p)\n", aboveProgName, val, &*val)
				program.ptrAbovePrograms = append(program.ptrAbovePrograms, val)
			}
		}
	}
	// fmt.Printf("\n")
	// for k, v := range programsMap {
	// 	fmt.Printf("%v (%p) %v\n", k, &*v, v)
	// }
	return programsMap
}

// ParseProgramLine parses a program line and returns corresponding properties (name of the program, id of the program, names of the programs immediately it)
func ParseProgramLine(programLine string) Program {
	progName := ""
	progWeight := -1
	var abovePrograms []string
	// qvymrle (166) -> daeabq, eskrg
	re := regexp.MustCompile(`(?P<name>[a-z]+) \((?P<weight>\d+)\)(?: -> (?P<above>[ ,a-z]+))?`)
	groupNames := re.SubexpNames()
	for _, match := range re.FindAllStringSubmatch(programLine, -1) {
		for groupIdx, group := range match {
			name := groupNames[groupIdx]
			if name == "name" {
				progName = group
			}
			if name == "weight" {
				progWeight, _ = strconv.Atoi(group)
			}
			if name == "above" && group != "" {
				abovePrograms = strings.Split(group, ", ")
			}
		}
	}
	return Program{progName, progWeight, abovePrograms, nil}
}

// Program ...
type Program struct {
	Name             string
	Weight           int
	abovePrograms    []string
	ptrAbovePrograms []*Program
}
