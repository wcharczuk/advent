package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/wcharczuk/advent/pkg/fileutil"
)

func main() {
	var board Board
	if err := fileutil.ReadByLines("../input", func(line string) error {
		row := make([]bool, len(line))
		for index, r := range line {
			if r == '#' {
				row[index] = true
			}
		}
		board = append(board, row)
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	var xmax, ymax, score, scoreMax int
	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[y]); y++ {
			if board[y][x] {
				if score = board.Score(x, y); score > scoreMax {
					xmax = x
					ymax = y
					scoreMax = score
				}
			}
		}
	}

	fmt.Printf("ANSWER: %d,%d = %d\n", xmax, ymax, scoreMax)
}

// Board is a 2 dimensional bool array.
type Board [][]bool

// String returns the board as a string.
func (b Board) String() string {
	var lines []string
	for _, row := range b {
		var line []rune
		for index, r := range row {
			if r {
				line[index] = '#'
			} else {
				line[index] = '.'
			}
		}
		lines = append(lines, string(line))
	}
	return strings.Join(lines, "\n")
}

// Get returns a value at a given coordinate.
func (b Board) Get(x, y int) bool {
	if y < len(b) {
		if x < len(b[y]) {
			return b[y][x]
		}
	}
	panic(fmt.Errorf("invalid coordinate: %d,%d", x, y))
}

// Height returns the board height.
func (b Board) Height() int {
	return len(b)
}

// Width returns the board width.
func (b Board) Width() int {
	if len(b) > 0 {
		return len(b[0])
	}
	return 0
}

// Score counts the other points reachable direclty from
// the given coordinates.
func (b Board) Score(x, y int) int {

}
