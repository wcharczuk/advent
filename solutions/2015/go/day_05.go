package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var VOWELS = map[string]bool{"a": true, "e": true, "i": true, "o": true, "u": true}

func checkWord(theString string) bool {
	isOk := true
	isOk = isOk && hasVowels(theString)
	isOk = isOk && hasRepeatedLetter(theString)
	isOk = isOk && doesntHaveBadWords(theString)

	return isOk
}

func hasVowels(theString string) bool {
	count := 0
	for _, c := range theString {
		if _, isVowel := VOWELS[string(c)]; isVowel {
			count = count + 1
		}
		if count == 3 {
			return true
		}
	}
	return false
}

func hasRepeatedLetter(theString string) bool {
	for i, c := range theString {
		if i < len(theString)-1 {
			if c == rune(theString[i+1]) {
				return true
			}
		}
	}
	return false
}

func doesntHaveBadWords(theString string) bool {
	badWords := []string{"ab", "cd", "pq", "xy"}
	for _, bad := range badWords {
		if strings.Contains(theString, bad) {
			return false
		}
	}
	return true
}

func main() {
	dataFile := "../testdata/day5"

	count := 0
	if f, err := os.Open(dataFile); err == nil {
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			word := scanner.Text()
			if checkWord(word) {
				count = count + 1
			}
		}
	}

	fmt.Printf("%d Ok Words\n", count)
}
