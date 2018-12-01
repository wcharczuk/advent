package main

import (
	"fmt"
	"strings"

	"github.com/blend/go-sdk/util"
)

const (
	ROWS  = 100
	COLS  = 100
	STEPS = 100
)

type Grid [ROWS][COLS]bool

func (g *Grid) TurnCornersOn() {
	g[0][0] = true
	g[ROWS-1][COLS-1] = true
	g[0][COLS-1] = true
	g[ROWS-1][0] = true
}

func (g *Grid) Step() {
	toTurnOn := [][]int{}
	toTurnOff := [][]int{}
	for row := 0; row < ROWS; row++ {
		for col := 0; col < COLS; col++ {
			result := g.Evaluate(row, col)
			if result {
				toTurnOn = append(toTurnOn, []int{row, col})
			} else {
				toTurnOff = append(toTurnOff, []int{row, col})
			}
		}
	}

	for _, pos := range toTurnOn {
		g[pos[0]][pos[1]] = true
	}

	for _, pos := range toTurnOff {
		g[pos[0]][pos[1]] = false
	}
}

func (g Grid) Evaluate(row, col int) bool {
	toEvaluate := [][]int{}

	toEvaluate = append(toEvaluate, []int{row - 1, col})
	toEvaluate = append(toEvaluate, []int{row - 1, col - 1})
	toEvaluate = append(toEvaluate, []int{row, col - 1})
	toEvaluate = append(toEvaluate, []int{row + 1, col - 1})
	toEvaluate = append(toEvaluate, []int{row + 1, col})
	toEvaluate = append(toEvaluate, []int{row + 1, col + 1})
	toEvaluate = append(toEvaluate, []int{row, col + 1})
	toEvaluate = append(toEvaluate, []int{row - 1, col + 1})

	countSurroundingOn := 0
	for _, pos := range toEvaluate {
		if pos[0] < 0 || pos[0] >= ROWS {
			continue
		}
		if pos[1] < 0 || pos[1] >= COLS {
			continue
		}
		if g[pos[0]][pos[1]] {
			countSurroundingOn++
		}
	}

	if g[row][col] {
		if countSurroundingOn == 2 || countSurroundingOn == 3 {
			return true
		}
		return false
	}
	if countSurroundingOn == 3 {
		return true
	}
	return false

}

func (g Grid) TotalOn() int {
	count := 0
	for row := 0; row < ROWS; row++ {
		for col := 0; col < COLS; col++ {
			if g[row][col] {
				count++
			}
		}
	}

	return count
}

func (g Grid) String() string {
	output := []string{}
	for row := 0; row < ROWS; row++ {
		rowText := ""
		for col := 0; col < COLS; col++ {
			if g[row][col] {
				rowText = rowText + "#"
			} else {
				rowText = rowText + "."
			}
		}
		output = append(output, rowText)
	}
	return strings.Join(output, "\n")
}

func main() {
	g := &Grid{}
	row := 0
	util.File.ReadByLines("../testdata/day18", func(line string) error {
		lineBytes := []byte(line)
		for col := 0; col < len(lineBytes); col++ {
			if string(lineBytes[col]) == "#" {
				g[row][col] = true
			} else {
				g[row][col] = false
			}
		}
		row++
		return nil
	})

	for x := 0; x < STEPS; x++ {
		g.Step()
	}

	fmt.Printf("After %d steps, %d are on\n", STEPS, g.TotalOn())
}
