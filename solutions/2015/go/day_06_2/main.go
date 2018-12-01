package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	ON     = iota
	OFF    = iota
	TOGGLE = iota
)

type point struct {
	Row, Col int
}

func (p point) String() string {
	return fmt.Sprintf("{%d,%d}", p.Row, p.Col)
}

type operation struct {
	Start  point
	End    point
	Effect int
}

func (op operation) String() string {
	effect := ""
	if op.Effect == ON {
		effect = "turn on"
	} else if op.Effect == OFF {
		effect = "turn off"
	} else {
		effect = "toggle"
	}

	return fmt.Sprintf("%s %s through %s", effect, op.Start.String(), op.End.String())
}

func main() {
	dataFile := "../testdata/day6"

	lights := [1000][1000]int{}

	if f, err := os.Open(dataFile); err == nil {
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			instruction := trim(scanner.Text())
			op := parseInstructions(strings.Split(instruction, " "))
			apply(op, &lights)
		}
	}

	count := countBrightness(lights)
	fmt.Printf("%d total brightness.\n", count)
}

func countBrightness(lights [1000][1000]int) int {
	count := 0
	for row := 0; row < 1000; row++ {
		for col := 0; col < 1000; col++ {
			count = count + lights[row][col]
		}
	}
	return count
}

func apply(op operation, lights *[1000][1000]int) {
	for row := op.Start.Row; row <= op.End.Row; row++ {
		for col := op.Start.Col; col <= op.End.Col; col++ {
			if op.Effect == ON {
				lights[row][col] = lights[row][col] + 1
			} else if op.Effect == OFF {
				if lights[row][col] == 0 {
					lights[row][col] = 0
				} else if lights[row][col] == 1 {
					lights[row][col] = 0
				} else {
					lights[row][col] = lights[row][col] - 1
				}
			} else if op.Effect == TOGGLE {
				lights[row][col] = lights[row][col] + 2
			}
		}
	}
}

func parseInstructions(instructionParts []string) operation {
	op := operation{}
	if instructionParts[0] == "turn" {
		if instructionParts[1] == "off" {
			op.Effect = OFF
		} else if instructionParts[1] == "on" {
			op.Effect = ON
		}

		op.Start = parsePoint(instructionParts[2])
		op.End = parsePoint(instructionParts[4])
	} else if instructionParts[0] == "toggle" {
		op.Effect = TOGGLE
		op.Start = parsePoint(instructionParts[1])
		op.End = parsePoint(instructionParts[3])
	}
	return op
}

func parsePoint(pointText string) point {
	pieces := strings.Split(pointText, ",")
	row, _ := strconv.Atoi(pieces[0])
	col, _ := strconv.Atoi(pieces[1])
	return point{row, col}
}

func trim(str string) string {
	return strings.Trim(str, " \t")
}
