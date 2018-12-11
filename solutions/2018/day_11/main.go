package main

import (
	"strconv"

	"github.com/wcharczuk/advent/pkg/array"
	"github.com/wcharczuk/advent/pkg/log"
)

const (
	width  = 300
	height = 300

	inputSerialNumber = 9221
)

/*
const (
	width  = 300
	height = 300

	inputSerialNumber = 18
)
*/

func main() {
	cells := array.TwoOfInt(300, 300)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			cells[x][y] = powerLevel(inputSerialNumber, x+1, y+1)
		}
	}

	var maxTotal, mx, my int
	for x := 0; x < width-2; x++ {
		for y := 0; y < height-2; y++ {
			total := total(cells, x, y)
			if maxTotal < total {
				maxTotal = total
				mx = x
				my = y
			}
		}
	}

	log.Solutionf("%d, %d", mx+1, my+1)
}

func total(cells [][]int, x, y int) int {
	var output int
	output += cells[x][y]
	output += cells[x+1][y]
	output += cells[x+2][y]

	output += cells[x][y+1]
	output += cells[x+1][y+1]
	output += cells[x+2][y+1]

	output += cells[x][y+2]
	output += cells[x+1][y+2]
	output += cells[x+2][y+2]

	return output
}

func powerLevel(serialNumber, x, y int) int {
	rackID := (x + 10)
	output := rackID * y
	output = output + serialNumber
	output = (output * rackID)

	chars := strconv.Itoa(output)
	if len(chars) < 3 {
		return 0
	}
	digit, _ := strconv.Atoi(string(chars[len(chars)-3]))
	return digit - 5
}
