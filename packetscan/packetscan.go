package packetscan

import (
	"regexp"
	"strconv"
	"strings"
)

const scannerPosition int = 0
const scannerDirection int = 1

// GoThroughTheFirewall returns the severity of the whole trip through the firewall
func GoThroughTheFirewall(content string) int {
	f := InitializeFirewall(content)
	severity, _ := goThroughTheFirewallWithDelay(content, 0, f)
	return severity
}

// goThroughTheFirewallWithDelay returns the severity of the whole trip through the firewall except if the delay variable is set in which case it returns the first sevrity encountered; it will also return the Firewall state before trying to through for later reuse
func goThroughTheFirewallWithDelay(content string, delay int, f Firewall) (int, Firewall) {
	// fmt.Printf("\n\nDelay %d\n", delay)

	currentPosition := -1
	if delay == 0 {
		currentPosition = -1
	} else {
		for f.picosecond < delay {
			f.NextPicosecond()
		}
	}

	scanners := make([]int, len(f.scanners))
	scannersDir := make([]int, len(f.scannersDir))
	copy(scanners, f.scanners)
	copy(scannersDir, f.scannersDir)
	fBak := Firewall{f.layers, scanners, scannersDir, f.picosecond}

	severity := 0
	for i := 0; i < len(f.scanners); i++ {
		currentPosition++
		// fmt.Printf("Move to %d (scanner %d at: %d)\n", currentPosition, i, f.scanners[i])
		// fmt.Printf("\tScanners (picosecond: %d) %v\n", f.picosecond, f.scanners)
		// Caught!!!
		if f.scanners[i] == 0 {
			severity += i * f.layers[i]
			// fmt.Printf("\tCaught! (%d)\n", severity)
			if delay > 0 {
				if severity == 0 {
					severity = 1
				}
				break
			}
		}
		f.NextPicosecond()
	}
	return severity, fBak
}

// GoSafelyThroughTheFirewall returns the number of picoseconds we need to wait before being able to go safely through the firewall
func GoSafelyThroughTheFirewall(content string) int {
	delay := 1
	f := InitializeFirewall(content)
	var severity int
	for {
		severity, f = goThroughTheFirewallWithDelay(content, delay, f)
		if severity == 0 {
			break
		}
		delay++
	}
	return delay
}

// ParseLine returns the progID and the communication programs IDs of a line of text
func parseLine(line string) (int, int) {
	layerID := -1
	depth := -1
	// 0 <-> 950, 1039
	re := regexp.MustCompile(`(?P<layerID>\d+): (?P<depth>\d+)`)
	groupNames := re.SubexpNames()
	for _, match := range re.FindAllStringSubmatch(line, -1) {
		for groupIdx, groupValue := range match {
			groupName := groupNames[groupIdx]
			switch {
			case groupName == "layerID":
				layerID, _ = strconv.Atoi(groupValue)
			case groupName == "depth":
				depth, _ = strconv.Atoi(groupValue)
			}
		}
	}
	return layerID, depth
}

// Firewall ...
type Firewall struct {
	layers      map[int]int
	scanners    []int // will record the current position of the scanners in each layer
	scannersDir []int // will record the current direction of the scanner in each layer (1 = going up; -1 going down)
	picosecond  int
}

// InitializeFirewall a firewall from the content of a string
func InitializeFirewall(content string) Firewall {
	var f Firewall
	lines := strings.Split(content, "\n")
	lastLayerID := 0
	f.layers = make(map[int]int)
	for _, line := range lines {
		layerID, depth := parseLine(line)
		f.layers[layerID] = depth
		lastLayerID = layerID
	}
	f.scanners = make([]int, lastLayerID+1)
	f.scannersDir = make([]int, lastLayerID+1)
	for i := 0; i <= lastLayerID; i++ {
		if _, ok := f.layers[i]; ok {
			f.scanners[i] = 0
			f.scannersDir[i] = 1
		} else {
			f.scanners[i] = -1
			f.scannersDir[i] = 0
		}
	}
	f.picosecond = 0
	return f
}

// NextPicosecond moves the scanners trhough the firewall for the next picosecond
func (f *Firewall) NextPicosecond() {
	// fmt.Printf("\nProcessing scanners from %d picoseconds", f.picosecond)
	f.picosecond++
	// fmt.Printf(" to %d picoseconds\n", f.picosecond)
	for i := range f.scanners {
		if depth, ok := f.layers[i]; ok {
			// fmt.Printf("\tProcessing scanner#%d (depth: %d)...", i, depth)
			// fmt.Printf("\t from %d", f.scanners[i])
			f.scanners[i] += f.scannersDir[i]
			if f.scanners[i] == depth-1 {
				f.scannersDir[i] = -1
			}
			if f.scanners[i] == 0 {
				f.scannersDir[i] = 1
			}
			// fmt.Printf(" to %d\n", f.scanners[i])
		}
	}
}
