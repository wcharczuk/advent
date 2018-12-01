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
	visits[point{0, 0}] = 1
	x, y := 0, 0
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
				switch current {
				case "^":
					x = x + 1
					break
				case ">":
					y = y + 1
					break
				case "v":
					x = x - 1
					break
				case "<":
					y = y - 1
					break
				}

				house := point{x, y}
				if count, hasVisited := visits[house]; hasVisited {
					visits[house] = count + 1
				} else {
					visits[house] = 1
				}
			}
		}
	}
	fmt.Printf("Santa visited %d houses.\n", len(visits))
}
