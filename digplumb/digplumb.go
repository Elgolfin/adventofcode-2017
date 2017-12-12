package digplumb

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/elgolfin/adventofcode-2017/sliceutil"
)

// GetAllProgGroups returns the number of program groups and the map of the program groups
func GetAllProgGroups(content string) (int, map[int]int) {
	progGroups := make(map[int]int)
	progArr := initializeProgArr(content)
	progHistory := make(map[int]bool)
	for progID := range progArr {
		if _, ok := progHistory[progID]; !ok {
			// fmt.Printf("Processing %d...\n", progID)
			currentHistory := make(map[int]bool)
			progGroups[progID], currentHistory = getProgGroupCount(progArr, progID)
			for k, v := range currentHistory {
				progHistory[k] = v
			}
		}
	}
	return len(progGroups), progGroups
}

// getProgGroupCount returns how many programs are in the group that contains the program ID
func getProgGroupCount(progArr [][]int, masterProgID int) (int, map[int]bool) {
	groupArr := []int{masterProgID}
	progHistory := make(map[int]bool)
	groupArr = processGroup(groupArr, masterProgID, progArr, progHistory)
	return len(groupArr), progHistory
}

func initializeProgArr(content string) [][]int {
	lines := strings.Split(content, "\n")
	progArr := make([][]int, len(lines))
	for _, line := range lines {
		progID, progGroup := parseLine(line)
		progArr[progID] = progGroup
	}
	return progArr
}

func processGroup(groupArr []int, masterProgID int, input [][]int, history map[int]bool) []int {
	for _, progID := range input[masterProgID] {
		groupArr = sliceutil.InsertInt(progID, groupArr)
		if _, ok := history[progID]; !ok {
			// fmt.Printf("Processing %d...\n", progID)
			history[progID] = true
			groupArr = processGroup(groupArr, progID, input, history)
		} else {
			// fmt.Printf("Not Processing %d because it has already been processed.\n", progID)
			history[progID] = true
		}
	}

	// fmt.Printf("%v\n", groupArr)
	return groupArr
}

// ParseLine returns the progID and the communication programs IDs of a line of text
func parseLine(line string) (int, []int) {
	progID := -1
	var progGroup []int
	// 0 <-> 950, 1039
	re := regexp.MustCompile(`(?P<progID>\d+) <-> (?P<progGroup>\d+(?:[ ,\d]+)?)`)
	groupNames := re.SubexpNames()
	for _, match := range re.FindAllStringSubmatch(line, -1) {
		for groupIdx, groupValue := range match {
			groupName := groupNames[groupIdx]
			switch {
			case groupName == "progID":
				progID, _ = strconv.Atoi(groupValue)
			case groupName == "progGroup":
				progGroup = sliceutil.Atoi(groupValue, ", ")
			}
		}
	}
	return progID, progGroup
}
