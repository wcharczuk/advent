package main

import (
	"github.com/wcharczuk/advent/pkg/fileutil"
	"github.com/wcharczuk/advent/pkg/log"
)

func diff(a, b string) int {
	if len(a) != len(b) {
		panic("unequal length strings")
	}

	var diff int
	for index := 0; index < len(a); index++ {
		if a[index] != b[index] {
			diff++
		}
	}
	return diff
}

func findFirstDifferenceIndex(a, b string) int {
	if len(a) != len(b) {
		panic("unequal length strings")
	}

	for index := 0; index < len(a); index++ {
		if a[index] != b[index] {
			return index
		}
	}
	return -1
}

func removeCharacterAt(value string, index int) string {
	if len(value) == 0 {
		return value
	}
	return value[:index] + value[index+1:]
}

func main() {
	var lines []string
	err := fileutil.ReadByLines("./input", func(line string) error {
		for _, previous := range lines {
			if delta := diff(line, previous); delta == 1 {
				index := findFirstDifferenceIndex(line, previous)
				log.Solution(removeCharacterAt(line, index))
			}
		}
		lines = append(lines, line)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal("no solution found")
}
