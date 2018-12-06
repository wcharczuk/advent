package main

import (
	"fmt"

	"github.com/wcharczuk/advent/pkg/fileutil"
	"github.com/wcharczuk/advent/pkg/log"
)

const (
	maxDistance = 10000
)

func main() {
	var points []Point
	err := fileutil.ReadByLines("./input", func(line string) error {
		var p Point
		_, err := fmt.Sscanf(line, "%d, %d", &p.X, &p.Y)
		if err != nil {
			return err
		}
		points = append(points, p)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	var mx, my int
	for _, p := range points {
		if mx < p.X {
			mx = p.X
		}
		if my < p.Y {
			my = p.Y
		}
	}
	mx++
	my++

	log.Context("debug").Printf("mx: %d my: %d", mx, my)

	// build the grid
	grid := make([][]int, my)
	for y := 0; y < my; y++ {
		grid[y] = make([]int, mx)
		for x := 0; x < mx; x++ {
			p2 := Point{X: x, Y: y}

			var totalDistance int
			for _, point := range points {
				totalDistance += point.Distance(p2)
			}

			if totalDistance < maxDistance {
				grid[y][x] = 1
			} else {
				grid[y][x] = -1
			}
		}
	}
	var total int
	for y := 0; y < my; y++ {
		for x := 0; x < mx; x++ {
			value := grid[y][x]
			if value == 1 {
				total++
			}
		}
	}

	log.Solution(total)
}

type Point struct {
	X, Y int
}

func (p Point) Distance(p2 Point) int {
	return abs(p.X-p2.X) + abs(p.Y-p2.Y)
}

func (p Point) String() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
