package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/wcharczuk/advent/pkg/fileutil"
)

func main() {
	var paths []path
	fileutil.ReadByLines("../input", func(line string) error {
		path := parsePath(strings.TrimSpace(line))
		paths = append(paths, path)
		return nil
	})

	points := findIntersections(paths...)
	if len(points) == 0 {
		fmt.Println("NO INTERSECTIONS")
		return
	}

	minDistance := points[0].Z
	for _, point := range points[1:] {
		distance := point.Z
		if minDistance > distance {
			minDistance = distance
		}
	}
	fmt.Printf("ANSWER: %d\n", minDistance)
}

func findIntersections(paths ...path) (output []point) {
	board := map[string]map[int]point{}

	for index, path := range paths {
		points := path.Expand()

		for _, p := range points {
			// if we haven't seen a given point, initialize it
			if _, ok := board[p.String()]; !ok {
				board[p.String()] = map[int]point{
					index: p,
				}
				continue
			}

			// if we've seen a given point
			previous := board[p.String()]
			// we haven't seen it for our index
			if _, ok := previous[index]; !ok {
				previous[index] = p
			}
		}
	}

	for _, position := range board {
		if len(position) > 1 {
			var x, y, z int
			for _, point := range position {
				x = point.X
				y = point.Y
				z = z + point.Z
			}
			output = append(output, point{
				X: x,
				Y: y,
				Z: z,
			})
		}
	}
	return
}

func parsePath(rawPath string) path {
	var output []segment
	for _, token := range strings.Split(rawPath, ",") {
		direction := token[0]
		distance, err := strconv.Atoi(token[1:])
		if err != nil {
			panic(err)
		}
		output = append(output, segment{
			Direction: string(direction),
			Distance:  distance,
		})
	}
	return path(output)
}

type point struct {
	X, Y, Z int // x==L|R, y==U|D, z==steps from start
}

func (p point) String() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

type path []segment

func (p path) Expand() (output []point) {
	var cursor point
	for _, segment := range p {
		for x := 0; x < segment.Distance; x++ {
			cursor.Z++
			switch segment.Direction {
			case "U":
				cursor.Y++
			case "D":
				cursor.Y--
			case "L":
				cursor.X--
			case "R":
				cursor.X++
			}
			output = append(output, cursor)
		}
	}
	return
}

func (p path) String() string {
	var pieces []string
	for _, seg := range p {
		pieces = append(pieces, seg.String())
	}
	return strings.Join(pieces, ",")
}

type segment struct {
	Direction string
	Distance  int
}

func (s segment) String() string {
	return s.Direction + strconv.Itoa(s.Distance)
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
