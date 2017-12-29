package electromoat

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// FindHeaviestBridges returns the weight of the heaviest bridge and the heaviest bridge that is the longest one
func FindHeaviestBridges(content string) (int, int) {
	availableComponents := LoadComponents(content)
	chains := BuildBridges("", make(map[string]int), 0, availableComponents)
	heaviestWeight := 0
	longestHeaviestBridge := 0
	longestBridge := 0
	test := make(map[int]int)
	for key := range chains {
		length, weight := GetBridgeStats(key)
		if weight > heaviestWeight {
			heaviestWeight = weight
			// key1 = key
		}
		test[length]++
		if length >= longestBridge {
			longestBridge = length
			if weight > longestHeaviestBridge {
				longestHeaviestBridge = weight
			}
		}
	}
	return heaviestWeight, longestHeaviestBridge
}

// GetBridgeStats returns the length and weight of a chain
func GetBridgeStats(components string) (int, int) {
	weight := 0
	length := 0

	f := func(c rune) bool {
		return !unicode.IsNumber(c)
	}

	for _, connector := range strings.FieldsFunc(components, f) {
		currentWeight, _ := strconv.Atoi(connector)
		weight += currentWeight
		length++
	}
	return length / 2, weight
}

// BuildBridges returns an array of all possible chains
func BuildBridges(chain string, chains map[string]int, port int, componentsMap map[int][]Component) map[string]int {
	for _, component := range componentsMap[port] {
		key := fmt.Sprintf("%s--%d,%d", chain, component.port1, component.port2)
		if !strings.Contains(chain, fmt.Sprintf("--%d,%d", component.port1, component.port2)) {
			chains[key] = chains[chain] + component.port1 + component.port2
			chains = BuildBridges(key, chains, component.getOtherPort(port), componentsMap)
		}
	}
	return chains
}

// LoadComponents returns an array of components from a string
func LoadComponents(content string) map[int][]Component {
	componentsStr := strings.Split(content, "\n")
	components := map[int][]Component{}
	for _, componentStr := range componentsStr {
		ports := strings.Split(componentStr, "/")
		port1, _ := strconv.Atoi(ports[0])
		port2, _ := strconv.Atoi(ports[1])
		component := Component{port1, port2}
		if _, ok := components[port1]; ok {
			components[port1] = append(components[port1], component)
		} else {
			components[port1] = []Component{component}
		}
		if port1 != port2 {
			if _, ok := components[port2]; ok {
				components[port2] = append(components[port2], component)
			} else {
				components[port2] = []Component{component}
			}
		}
	}
	return components
}

func (component Component) getOtherPort(port int) int {
	if component.port1 == port {
		return component.port2
	}
	return component.port1
}

// Component ...
type Component struct {
	port1 int
	port2 int
}
