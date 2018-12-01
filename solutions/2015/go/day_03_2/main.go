package main

import (
	"fmt"
	"io"
	"os"
)

type point struct {
	X, Y int
}

func (p point) String() string {
	return fmt.Sprintf("{%d,%d}", p.X, p.Y)
}

func main() {
	dataFile := "../testdata/day3"

	visits := map[point]int{}
	visits[point{0, 0}] = 2

	santa_x, santa_y := 0, 0
	robo_x, robo_y := 0, 0
	instructions := 0

	if f, err := os.Open(dataFile); err == nil {
		defer f.Close()

		chunk := make([]byte, 32)
		for {
			readBytes, err := f.Read(chunk)
			if err == io.EOF {
				break
			}
			readData := chunk[:readBytes]
			for _, b := range readData {
				current := string(b)

				var house point
				if instructions%2 == 0 {
					switch current {
					case "^":
						santa_x = santa_x + 1
						break
					case ">":
						santa_y = santa_y + 1
						break
					case "v":
						santa_x = santa_x - 1
						break
					case "<":
						santa_y = santa_y - 1
						break
					}
					house = point{santa_x, santa_y}
				} else {
					switch current {
					case "^":
						robo_x = robo_x + 1
						break
					case ">":
						robo_y = robo_y + 1
						break
					case "v":
						robo_x = robo_x - 1
						break
					case "<":
						robo_y = robo_y - 1
						break
					}
					house = point{robo_x, robo_y}
				}

				if count, hasVisited := visits[house]; hasVisited {
					visits[house] = count + 1
				} else {
					visits[house] = 1
				}

				instructions = instructions + 1
			}
		}
	}
	fmt.Printf("Santa visited %d houses.\n", len(visits))
}
