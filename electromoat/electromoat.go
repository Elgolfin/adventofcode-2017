package electromoat

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// FindHeaviestBridge returns the weight of the heaviest bridge
func FindHeaviestBridge(content string) int {
	availableComponents := LoadComponents(content)
	chains := BuildChains("", make(map[string]int), 0, availableComponents)
	heaviestWeight := 0
	for key := range chains {
		_, weight := GetBridgeStats(key)
		if weight > heaviestWeight {
			heaviestWeight = weight
		}
	}
	return heaviestWeight
}

// CalculateWeight returns the weight of a chain
func CalculateWeight(components []Component) int {
	weight := 0
	for _, component := range components {
		weight += component.port1 + component.port2
	}
	return weight
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
	return length, weight
}

// BuildChains returns an array of all possible chains
func BuildChains(chain string, chains map[string]int, port int, componentsMap map[int][]Component) map[string]int {
	if len(componentsMap[port]) > 0 {
		for _, component := range componentsMap[port] {
			key := strings.TrimSpace(fmt.Sprintf("%s--%d,%d", chain, component.port1, component.port2))
			//if !IsComponentIn(component, chains[chain]) {
			if !strings.Contains(chain, fmt.Sprintf("%d,%d", component.port1, component.port2)) {
				chains[key] = chains[chain] + component.port1 + component.port2
				chains = BuildChains(key, chains, component.getOtherPort(port), componentsMap)
			}
		}
	}
	return chains
}

// IsComponentIn returns
func IsComponentIn(component Component, components []Component) bool {
	for _, c := range components {
		if c.port1 == component.port1 && c.port2 == component.port2 {
			return true
		}
	}
	return false
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
