package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func checkWord(theString string) bool {
	return hasRepeatedLetterInGap(theString) && hasRepeatedPair(theString)
}

func hasRepeatedPair(theString string) bool {
	for i := 0; i < len(theString)-3; i++ {
		first := rune(theString[i])
		second := rune(theString[i+1])

		for j := i + 2; j < len(theString)-1; j++ {
			testFirst := rune(theString[j])
			testSecond := rune(theString[j+1])

			if first == testFirst && second == testSecond {
				return true
			}
		}
	}
	return false
}

func hasRepeatedLetterInGap(theString string) bool {
	for i := 0; i < len(theString)-2; i++ {
		c := theString[i]
		c2 := theString[i+2]

		if c == c2 {
			return true
		}
	}
	return false
}

func main() {
	dataFile := "../testdata/day5"

	count := 0
	if f, err := os.Open(dataFile); err == nil {
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			word := strings.Trim(scanner.Text(), " \t\n")
			if checkWord(word) {
				count = count + 1
			}
		}
	}

	fmt.Printf("%d nice words\n", count)
}
