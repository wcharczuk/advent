package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/wcharczuk/advent/pkg/fileutil"
)

func main() {
	var orbits []Orbit
	if err := fileutil.ReadByLines("../input", func(line string) error {
		pieces := strings.Split(strings.TrimSpace(line), ")")
		orbits = append(orbits, Orbit{
			Center:  pieces[0],
			Orbiter: pieces[1],
		})
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	planets := BuildPlanets(orbits...)
	you := planets["YOU"]
	santa := planets["SAN"]
	if you == nil || santa == nil {
		log.Fatal(fmt.Errorf("Didn't find you or santa"))
	}

	path := FindSanta(you, santa)
	fmt.Printf("ANSWER: %d\n", len(path))
}

// FindSanta finds the path from to the other node.
func FindSanta(from, to *Planet) []string {
	return findSantaImpl(from, to, nil)
}

func findSantaImpl(from, to *Planet, path []string) []string {
	if from.Center != nil && to.Center != nil && from.Center.Name == to.Center.Name {
		return path
	}
	if from.Checksum() > to.Checksum() {
		return findSantaImpl(from.Center, to, append(path, from.Name))
	}
	if from.Checksum() < to.Checksum() {
		return findSantaImpl(from, to.Center, append(path, to.Name))
	}
	return findSantaImpl(from.Center, to.Center, append(path, from.Name, to.Name))
}

// Orbit is an orbit listing.
type Orbit struct {
	Center  string
	Orbiter string
}

// BuildPlanets builds the planet list.
func BuildPlanets(orbits ...Orbit) (lookup map[string]*Planet) {
	lookup = map[string]*Planet{}
	for _, orbit := range orbits {
		if _, ok := lookup[orbit.Center]; !ok {
			lookup[orbit.Center] = &Planet{
				Name: orbit.Center,
			}
		}
		if _, ok := lookup[orbit.Orbiter]; !ok {
			lookup[orbit.Orbiter] = &Planet{
				Name: orbit.Orbiter,
			}
		}
	}
	for _, orbit := range orbits {
		lookup[orbit.Center].Orbiters = append(lookup[orbit.Center].Orbiters, lookup[orbit.Orbiter])
		lookup[orbit.Orbiter].Center = lookup[orbit.Center]
	}
	return
}

// Planet is a final graph of the planets.
type Planet struct {
	Name     string
	Center   *Planet
	Orbiters []*Planet
}

// DFS walks the planet graph depth first.
func (p *Planet) DFS(visitor func(*Planet)) {
	visitor(p)
	if p.Center != nil {
		p.Center.DFS(visitor)
	}
	for _, child := range p.Orbiters {
		child.DFS(visitor)
	}
}

func (p *Planet) String() string {
	var names []string
	p.DFS(func(p2 *Planet) { names = append(names, fmt.Sprintf("%s(%d)", p2.Name, p2.Checksum())) })
	return strings.Join(names, ", ")
}

// Checksum computes the number of orbits for a given planet.
func (p *Planet) Checksum() int {
	if p.Center == nil {
		return 0
	}
	return 1 + p.Center.Checksum()
}
