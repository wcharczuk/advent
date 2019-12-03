package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/wcharczuk/advent/pkg/fileutil"
)

func main() {
	var paths []Path
	fileutil.ReadByLines("../input", func(line string) error {
		path := ParsePath(strings.TrimSpace(line))
		paths = append(paths, path)
		return nil
	})

	points := FindIntersections(paths...)
	if len(points) == 0 {
		fmt.Println("NO INTERSECTIONS")
		return
	}

	minDistance := points[0].Distance()
	for _, point := range points[1:] {
		distance := point.Distance()
		if minDistance > distance {
			minDistance = distance
		}
	}

	fmt.Printf("ANSWER: %d\n", minDistance)
}

func FindIntersections(paths ...Path) (output []Point) {
	seen := map[string]int{}
	for index, path := range paths {
		points := path.Expand()
		for _, p := range points {
			if seenIndex, ok := seen[p.String()]; ok {
				if seenIndex != index {
					output = append(output, p)
				}
			} else {
				seen[p.String()] = index
			}
		}
	}
	return
}

func ParsePath(path string) Path {
	var output []Segment
	for _, token := range strings.Split(path, ",") {
		direction := token[0]
		distance, err := strconv.Atoi(token[1:])
		if err != nil {
			panic(err)
		}
		output = append(output, Segment{
			Direction: string(direction),
			Distance:  distance,
		})
	}
	return Path(output)
}

type Point struct {
	X, Y int
}

func (p Point) String() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

func (p Point) Distance() int {
	return Abs(p.X) + Abs(p.Y)
}

type Path []Segment

func (p Path) Expand() (path []Point) {
	var cursor Point
	for _, segment := range p {
		for x := 0; x < segment.Distance; x++ {
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
			path = append(path, cursor)
		}
	}
	return
}

func (p Path) String() string {
	var pieces []string
	for _, seg := range p {
		pieces = append(pieces, seg.String())
	}
	return strings.Join(pieces, ",")
}

type Segment struct {
	Direction string
	Distance  int
}

func (s Segment) String() string {
	return s.Direction + strconv.Itoa(s.Distance)
}

func Abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
