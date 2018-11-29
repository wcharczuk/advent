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

	lights := [1000][1000]bool{}

	if f, err := os.Open(dataFile); err == nil {
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			instruction := trim(scanner.Text())
			op := parseInstructions(strings.Split(instruction, " "))
			apply(op, &lights)
		}
	}

	count := countOn(lights)
	fmt.Printf("%d lights are on.\n", count)
}

func countOn(lights [1000][1000]bool) int {
	count := 0
	for row := 0; row < 1000; row++ {
		for col := 0; col < 1000; col++ {
			if lights[row][col] {
				count = count + 1
			}
		}
	}
	return count
}

func apply(op operation, lights *[1000][1000]bool) {
	for row := op.Start.Row; row <= op.End.Row; row++ {
		for col := op.Start.Col; col <= op.End.Col; col++ {
			if op.Effect == ON {
				lights[row][col] = true
			} else if op.Effect == OFF {
				lights[row][col] = false
			} else if op.Effect == TOGGLE {
				lights[row][col] = !lights[row][col]
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
