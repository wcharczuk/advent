package main

import (
	"fmt"
	"sort"

	"github.com/blend/go-sdk/util"
)

var containers = []int{
	43,
	3,
	4,
	10,
	21,
	44,
	4,
	6,
	47,
	41,
	34,
	17,
	17,
	44,
	36,
	31,
	46,
	9,
	27,
	38,
}

func main() {
	targetValue := 150
	sort.Sort(sort.Reverse(sort.IntSlice(containers)))
	possibles := combinationsToTarget(targetValue, containers)
	fmt.Printf("Number of Possibles: %d\n", len(possibles))
}

func combinationsToTarget(target int, values []int) [][]int {
	possibleValues := util.Math.PowOfInt(2, uint(len(values)))

	output := [][]int{}

	for of := possibleValues; of > 0; of-- {
		row := []int{}
		for i := 0; i < len(values); i++ {
			y := 1 << uint(i)
			if y&of == 0 && y != of {
				row = append(row, values[i])
			}
		}
		if len(row) > 0 && util.Math.SumOfInt(row) == target {
			output = append(output, row)
		}
	}
	return output
}
