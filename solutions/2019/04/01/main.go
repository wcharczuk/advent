package main

import (
	"fmt"
	"strconv"
)

// it is a six-digit number
// value is within the range given in the puzzle input
// two adjacent digits are the same (and only two)
// going from left to right, the digits never decrease

// how many passwords within the range given in your puzzle input match this criteria?

func main() {
	result := process(264793, 803935, evaluator)
	fmt.Printf("answer is: %d\n", result)
}

func process(start, stop int, evaluator func(string) bool) (output int64) {
	for x := start; x < stop; x++ {
		if evaluator(strconv.Itoa(x)) {
			output++
		}
	}
	return
}

func evaluator(value string) bool {
	if len(value) != 6 {
		return false
	}
	if !twoAdjacentDigitsAreTheSame(value) {
		return false
	}
	if !digitsIncrease(value) {
		return false
	}
	return true
}

func twoAdjacentDigitsAreTheSame(value string) bool {
	return len(findRepeats(value)) >= 1 // tweak this
}

func findRepeats(value string) [][]rune {
	runes := []rune(value)

	var repeats [][]rune
	var working []rune

	var current, next rune

	for x := 0; x < len(runes)-1; x++ {
		current = runes[x]
		next = runes[x+1]
		if current == next {
			working = append(working, current)
		} else if len(working) > 0 {
			repeats = append(repeats, working)
			working = nil
		}
	}
	if len(working) > 0 {
		repeats = append(repeats, working)
		working = nil
	}
	return repeats
}

func digitsIncrease(value string) bool {
	runes := []rune(value)
	previous := runes[0]
	var current rune
	for x := 1; x < len(runes); x++ {
		current = runes[x]
		if current < previous {
			return false
		}
		previous = current
	}
	return true
}
