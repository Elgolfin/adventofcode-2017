package partswarm

import (
	"math"
	"regexp"
	"strings"

	"github.com/elgolfin/adventofcode-2017/sliceutil"
)

// GetClosestParticleToOrigin returns the particle that will stay closest to position <0,0,0> in the long term
func GetClosestParticleToOrigin(content string) int {
	count := 0
	closestParticleID := -1
	previousClosestParticleID := -1
	closestParticleDistanceFromOrigin := -1
	particlesStr := strings.Split(content, "\n")
	particles := make([]Particle, len(particlesStr))

	for ID, particleStr := range particlesStr {
		particles[ID] = InitializeParticle(particleStr)
	}

	for tick := 0; ; tick++ {
		for particleID := range particles {
			particle := &particles[particleID]
			particleDistanceFromOrigin := particle.GetManhattanDistanceFromOrigin()
			if particleID == 0 {
				closestParticleID = particleID
				closestParticleDistanceFromOrigin = particleDistanceFromOrigin
			}
			if particleDistanceFromOrigin < closestParticleDistanceFromOrigin {
				closestParticleID = particleID
				closestParticleDistanceFromOrigin = particleDistanceFromOrigin
			}
			particle.Update()
		}
		if previousClosestParticleID == closestParticleID {
			count++
		} else {
			previousClosestParticleID = closestParticleID
			count = 1
		}
		if count >= 1000 {
			break
		}
	}
	return closestParticleID
}

// InitializeParticle returns a particle struct from an input string
func InitializeParticle(content string) Particle {
	pStr, vStr, aStr := parseLine(content)
	p, v, a := InitializeCoordinates(pStr), InitializeCoordinates(vStr), InitializeCoordinates(aStr)
	return Particle{p, v, a}
}

func parseLine(line string) (string, string, string) {
	p := ""
	v := ""
	a := ""
	re := regexp.MustCompile(`p=<(?P<p>-?\d+,-?\d+,-?\d+)>, v=<(?P<v>-?\d+,-?\d+,-?\d+)>, a=<(?P<a>-?\d+,-?\d+,-?\d+)>`)
	groupNames := re.SubexpNames()
	for _, match := range re.FindAllStringSubmatch(line, -1) {
		for groupIdx, groupValue := range match {
			groupName := groupNames[groupIdx]
			switch {
			case groupName == "p":
				p = groupValue
			case groupName == "v":
				v = groupValue
			case groupName == "a":
				a = groupValue
			}
		}
	}
	return p, v, a
}

// InitializeCoordinates returns a coordinates struct from an input string
func InitializeCoordinates(content string) Coordinates {
	axis := sliceutil.Atoi(content, ",")
	return Coordinates{axis[0], axis[1], axis[2]}
}

// Update returns the new position of the particle after the next tick
func (p *Particle) Update() {
	p.velocity.x += p.accleration.x
	p.velocity.y += p.accleration.y
	p.velocity.z += p.accleration.z
	p.position.x += p.velocity.x
	p.position.y += p.velocity.y
	p.position.z += p.velocity.z
}

// GetManhattanDistanceFromOrigin returns the manhataan distance of the particle from the origin <0,0,0>
func (p Particle) GetManhattanDistanceFromOrigin() int {
	return int(math.Abs(float64(p.position.x)) + math.Abs(float64(p.position.y)) + math.Abs(float64(p.position.z)))
}

// Particle ...
type Particle struct {
	position    Coordinates
	velocity    Coordinates
	accleration Coordinates
}

// Coordinates ...
type Coordinates struct {
	x int
	y int
	z int
}
