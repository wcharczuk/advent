package main

import (
	"fmt"
	"strconv"
	"sync"

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

	results := make([]Result, 300)
	wg := sync.WaitGroup{}
	wg.Add(300)
	for size := 1; size < 300; size++ {
		go func(size int) {
			defer wg.Done()

			var maxTotal, mx, my int
			for x := 0; x < width-(size-1); x++ {
				for y := 0; y < height-(size-1); y++ {
					total := total(cells, x, y, size)
					if maxTotal < total {
						maxTotal = total
						mx = x
						my = y
					}
				}
			}
			results[size] = Result{Total: maxTotal, Size: size, X: mx, Y: my}
		}(size)
	}
	wg.Wait()

	println("tabulating results")
	var maxSize, maxMetaTotal, mmx, mmy int
	for x := 0; x < 300; x++ {
		result := results[x]
		if maxMetaTotal < result.Total {
			maxMetaTotal = result.Total
			maxSize = result.Size
			mmx = result.X + 1
			mmy = result.Y + 1
		}
	}

	log.Solutionf("%d,%d,%d", mmx, mmy, maxSize)
}

type Result struct {
	Total int
	Size  int
	X     int
	Y     int
}

func (r Result) String() string {
	return fmt.Sprintf("%d,%d,%d", r.X, r.Y, r.Size)
}

func total(cells [][]int, x, y int, size int) int {
	var output int
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			output += cells[x+i][y+j]
		}
	}
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
