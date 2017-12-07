package recurscircus

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/elgolfin/adventofcode-2017/sliceutil"
)

// GetRootProgram returns the root program
func GetRootProgram(content string) string {
	rootProgram := ""

	// Construct the map (only keep programs that have above programs)
	programs := strings.Split(content, "\n")
	programsMap := make(map[string][]string)
	for _, program := range programs {
		name, _, abovePrograms := ParseProgramLine(program)
		if len(abovePrograms) > 0 {
			programsMap[name] = abovePrograms
		}
	}

	// Clean-up the map by removing each program whose all the above programs are not in the map
	// Continue until there is just one remaining program (the root one)
	for len(programsMap) > 1 {
		programsToBeRemoved := make([]string, 0)
		for progName, abovePrograms := range programsMap {
			count := len(abovePrograms)
			for _, aboveProgram := range abovePrograms {
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
	return rootProgram
}

// ParseProgramLine parses a program line and returns corresponding properties (name of the program, id of the program, names of the programs immediately it)
func ParseProgramLine(programLine string) (string, int, []string) {
	progName := ""
	progWeight := -1
	var abovePrograms []string
	// qvymrle (166) -> daeabq, eskrg
	re := regexp.MustCompile(`(?P<name>[a-z]+) \((?P<id>\d+)\)(?: -> (?P<above>[ ,a-z]+))?`)
	groupNames := re.SubexpNames()
	for _, match := range re.FindAllStringSubmatch(programLine, -1) {
		for groupIdx, group := range match {
			name := groupNames[groupIdx]
			if name == "name" {
				progName = group
			}
			if name == "id" {
				progWeight, _ = strconv.Atoi(group)
			}
			if name == "above" && group != "" {
				abovePrograms = strings.Split(group, ",")
			}
		}
	}
	return progName, progWeight, abovePrograms
}
