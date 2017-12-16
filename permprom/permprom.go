package permprom

import (
	"regexp"
	"strconv"
	"strings"
)

// BillionDance returns the order of the programs after one billion dances
func BillionDance(content string, n int) string {
	lines := strings.Split(content, ",")
	moves := make([]Move, len(lines))
	for _, line := range lines {
		moves = append(moves, parseLine(line))
	}
	firstDance := ""
	historyDances := []string{}
	p := InitializePrograms(n)
	// Only perform the loop until we find two identical programs standing after executing many dances
	// It will mean that every n dances, the dances ends with the same results (so, there is no need to execute 1,000,000,000 times the loop)
	for i := 1; i <= 10000000000; i++ {
		p.Dance(moves)
		if i == 1 {
			firstDance = string(p.progs)
			historyDances = append(historyDances, firstDance)
		} else {
			currentDance := string(p.progs)
			if currentDance == firstDance {
				break
			}
			historyDances = append(historyDances, currentDance)
		}
	}

	// Find the last dance standing
	lastDance := 10000000000 % len(historyDances)
	// If the remainder equals 0, it means the last dance standing is the last dance that has been recorded in the history
	// Otherwise we need to substract one to the remainder to find the right index in the array of dances
	if lastDance == 0 {
		lastDance = len(historyDances) - 1
	} else {
		lastDance--
	}
	return historyDances[lastDance]
}

// Dance executes a series of dance moves with the sepcified programs
func (p *Programs) Dance(moves []Move) {
	for _, move := range moves {
		switch {
		case move.name == "Spin":
			// fmt.Printf("Spin %d\n", move.size)
			p.Spin(move.size)
		case move.name == "Exchange":
			// fmt.Printf("Exchange %d with %d\n", move.pos1, move.pos2)
			p.Exchange(move.pos1, move.pos2)
		case move.name == "Partner":
			// fmt.Printf("Partner %v with %v\n", string(move.prog1), string(move.prog2))
			p.Partner(move.prog1, move.prog2)
		}
	}
}

// Dance returns the order of the programs after the dance
func Dance(content string, n int) string {
	lines := strings.Split(content, ",")
	moves := make([]Move, len(lines))
	for _, line := range lines {
		moves = append(moves, parseLine(line))
	}
	p := InitializePrograms(n)
	p.Dance(moves)
	return string(p.progs)
}

// InitializePrograms returns a Programs struct properly initialized
func InitializePrograms(n int) Programs {
	var progs []rune
	for i := 0; i < n; i++ {
		progs = append(progs, rune(97+i))
	}
	return Programs{progs}
}

// Spin makes X programs move from the end to the front, but maintain their order otherwise
func (p *Programs) Spin(x int) {
	length := len(p.progs)
	toStart := p.progs[length-x : length]
	toEnd := p.progs[0 : length-x]
	newProgs := append(toStart, toEnd...)
	copy(p.progs, newProgs)
}

// Exchange makes the programs at positions A and B swap places
func (p *Programs) Exchange(pos1 int, pos2 int) {
	p.progs[pos1], p.progs[pos2] = p.progs[pos2], p.progs[pos1]
}

// Partner makes the programs named A and B swap places
func (p *Programs) Partner(p1 rune, p2 rune) {
	pos1 := -1
	pos2 := -1
	for i, r := range p.progs {
		if r == p1 {
			pos1 = i
		}
		if r == p2 {
			pos2 = i
		}
		if (pos1 > 0 || pos2 > 0) && pos1*pos2 > 0 {
			break
		}
	}
	p.Exchange(pos1, pos2)
}

func parseLine(line string) Move {
	name := ""
	size := 0
	pos1 := 0
	pos2 := 0
	var prog1 rune
	var prog2 rune
	re := regexp.MustCompile(`s(?P<size>\d+)|x(?P<pos1>\d+)\/(?P<pos2>\d+)|p(?P<prog1>[a-z])\/(?P<prog2>[a-z])`)
	groupNames := re.SubexpNames()
	for _, match := range re.FindAllStringSubmatch(line, -1) {
		for groupIdx, groupValue := range match {
			groupName := groupNames[groupIdx]
			switch {
			case groupName == "size" && groupValue != "":
				name = "Spin"
				size, _ = strconv.Atoi(groupValue)
			case groupName == "pos1" && groupValue != "":
				name = "Exchange"
				pos1, _ = strconv.Atoi(groupValue)
			case groupName == "pos2" && groupValue != "":
				pos2, _ = strconv.Atoi(groupValue)
			case groupName == "prog1" && groupValue != "":
				name = "Partner"
				prog1 = rune(groupValue[0])
			case groupName == "prog2" && groupValue != "":
				prog2 = rune(groupValue[0])
			}
		}
	}
	m := Move{name, size, pos1, pos2, prog1, prog2}
	return m
}

// Programs represents the list of prgrams who are participating in the dance
type Programs struct {
	progs []rune
}

// Move represents a Dance move
type Move struct {
	name  string
	size  int
	pos1  int
	pos2  int
	prog1 rune
	prog2 rune
}
