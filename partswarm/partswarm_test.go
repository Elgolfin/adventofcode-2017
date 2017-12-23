package partswarm

import (
	"testing"
)

func TestResolveCollisions(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{

		{`p=<-6,0,0>, v=<3,0,0>, a=<0,0,0>
p=<-4,0,0>, v=<2,0,0>, a=<0,0,0>
p=<-2,0,0>, v=<1,0,0>, a=<0,0,0>
p=<3,0,0>, v=<-1,0,0>, a=<0,0,0>`, 1,
		},
	}
	for _, c := range cases {
		got := ResolveCollisions(c.in)
		if got != c.want {
			t.Errorf("Expected ResolveCollisions(%v) to return %v, got %v", c.in, c.want, got)
		}
	}
}

func TestGetClosestParticleToOrigin(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{

		{`p=<3,0,0>, v=<2,0,0>, a=<-1,0,0>
			p=<4,0,0>, v=<0,0,0>, a=<-2,0,0>`, 0,
		},
		{`p=<4,0,0>, v=<0,0,0>, a=<-2,0,0>
			p=<3,0,0>, v=<2,0,0>, a=<-1,0,0>`, 1,
		},
	}
	for _, c := range cases {
		got := GetClosestParticleToOrigin(c.in)
		if got != c.want {
			t.Errorf("Expected GetClosestParticleToOrigin(%v) to return %v, got %v", c.in, c.want, got)
		}
	}
}

func TestUpdate(t *testing.T) {
	cases := []struct {
		in   string
		want Particle
	}{
		{`p=<1609,-863,-779>, v=<-15,54,-69>, a=<-10,0,14>`, Particle{Coordinates{1584, -809, -834}, Coordinates{-25, 54, -55}, Coordinates{-10, 0, 14}}},
	}
	for _, c := range cases {
		got := InitializeParticle(c.in)
		got.Update()
		if got != c.want {
			t.Errorf("Expected %v.Update() to return %v, got %v", c.in, c.want, got)
		}
	}
}
func TestGetManhattanDistanceFromOrigin(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{`p=<1609,-863,-779>, v=<-15,54,-69>, a=<-10,0,14>`, 3251},
		{`p=<3,0,0>, v=<2,0,0>, a=<-1,0,0>`, 3},
		{`p=<4,0,0>, v=<0,0,0>, a=<-2,0,0>`, 4},
		{`p=<-238,111,134>, v=<-3,7,-5>, a=<1,-1,0>`, 483},
	}
	for _, c := range cases {
		p := InitializeParticle(c.in)
		got := p.GetManhattanDistanceFromOrigin()
		if got != c.want {
			t.Errorf("Expected %v.TestGetManhattanDistanceFromOrigin() to return %v, got %v", c.in, c.want, got)
		}
	}
}

func TestInitializeParticle(t *testing.T) {
	cases := []struct {
		in   string
		want Particle
	}{
		{`p=<1609,-863,-779>, v=<-15,54,-69>, a=<-10,0,14>`, Particle{Coordinates{1609, -863, -779}, Coordinates{-15, 54, -69}, Coordinates{-10, 0, 14}}},
	}
	for _, c := range cases {
		got := InitializeParticle(c.in)
		if got != c.want {
			t.Errorf("Expected -10,0,14(%q) to return %d, got %d", c.in, c.want, got)
		}
	}
}
