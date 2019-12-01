package main

import (
	"fmt"
	"log"
	"math"

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

// ComputeFuel computes the fuel for a given mass.
func ComputeFuel(mass float64) float64 {
	return math.Floor(mass/3.0) - 2.0
}
