package main

import (
	"fmt"
	"math"

	"github.com/wcharczuk/advent/pkg/collections"

	"github.com/wcharczuk/advent/pkg/fileutil"
	"github.com/wcharczuk/advent/pkg/log"
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

			distances := make([]int, len(points))
			minDistance := int(math.MaxInt64)

			var minID int
			for id, p := range points {
				distance := p.Distance(p2)
				distances[id] = distance
				if distance < minDistance {
					minDistance = distance
					minID = id
				}
			}

			var existsTie bool
			for id, distance := range distances {
				if id != minID && distance == minDistance {
					existsTie = true
					break
				}
			}

			if existsTie {
				grid[y][x] = -1
			} else {
				grid[y][x] = minID
			}
		}
	}

	// find infinite zones
	infinite := collections.NewSetOfInt()
	for y := 0; y < my; y++ {
		for x := 0; x < mx; x++ {
			value := grid[y][x]
			if y == 0 || x == 0 || x == mx-1 || y == my-1 {
				if value != -1 {
					infinite.Add(value)
				}
			}
		}
	}

	counts := map[int]int{}
	for y := 0; y < my; y++ {
		for x := 0; x < mx; x++ {
			value := grid[y][x]
			if value != -1 && !infinite.Contains(value) {
				counts[value] = counts[value] + 1
			}
		}
	}

	var maxCount int
	for _, count := range counts {
		if count > maxCount {
			maxCount = count
		}
	}

	log.Solution(maxCount)
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
