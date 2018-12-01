package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dataFile := "../testdata/day2"

	total := 0
	if f, err := os.Open(dataFile); err == nil {
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			pieces := strings.Split(line, "x")

			sides := make([]int, 3)
			for i := 0; i < 3; i++ {
				if parsed, parsedErr := strconv.Atoi(pieces[i]); parsedErr == nil {
					sides[i] = parsed
				}
			}

			l, w, h := sides[0], sides[1], sides[2]
			side1 := l * w
			side2 := w * h
			side3 := h * l
			total = total + 2*side1 + 2*side2 + 2*side3

			total = total + min(side1, side2, side3)
		}
	}

	fmt.Printf("Total Feet Requried : %d\n", total)
}

func min(args ...int) int {
	minValue := 1 << 31
	for _, v := range args {
		if v < minValue {
			minValue = v
		}
	}
	return minValue
}
