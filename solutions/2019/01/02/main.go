package main

import (
	"fmt"
	"log"
	"math"

	"github.com/wcharczuk/advent/pkg/dbg"
	"github.com/wcharczuk/advent/pkg/fileutil"
)

func main() {
	weights, err := fileutil.ReadFloats("../input")
	if err != nil {
		log.Fatal(err)
	}

	var totalFuel float64
	for _, weight := range weights {
		totalFuel += ComputeFuel(weight)
	}
	fmt.Printf("ANSWER: %0.4f", totalFuel)
}

// ComputeFuel computes the fuel required to lift the mass
// and the fuel required to lift that mass.
func ComputeFuel(mass float64) (output float64) {
	for {
		dbg.Spew(mass)
		fuel := ComputeFuelForMass(mass)
		if fuel == 0 {
			return
		}
		output += fuel
		mass = fuel
	}
}

// ComputeFuelForMass computes the fuel for a given mass.
func ComputeFuelForMass(mass float64) float64 {
	value := math.Floor(mass/3.0) - 2.0
	if value < 0 {
		return 0
	}
	return value
}
