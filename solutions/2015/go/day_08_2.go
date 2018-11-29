package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type testCase struct {
	Value  string
	Memory int
	Code   int
}

func main() {
	dataFile := "../testdata/day8"

	testCases := []testCase{
		{`""`, 6, 2},
		{`"abc"`, 9, 5},
		{`"aaa\"aaa"`, 16, 10},
		{`"\x27"`, 11, 6},
	}

	allPassed := true
	for _, test := range testCases {
		memory, code := process(test.Value)
		passed := code == test.Code && memory == test.Memory

		fmt.Println("Test", test.Value, "Code:", code, test.Code, "Memory:", memory, test.Memory)
		allPassed = allPassed && passed
	}

	if !allPassed {
		fmt.Println("Test cases failed, aborting!")
		os.Exit(1)
	}

	codeCount := 0
	memoryCount := 0
	if f, err := os.Open(dataFile); err == nil {
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			word := scanner.Text()
			memory, code := process(word)
			memoryCount += memory
			codeCount += code
		}
	}

	fmt.Println("Code:", codeCount, "Memory:", memoryCount, "Diff:", memoryCount-codeCount)
}

func process(input string) (int, int) {
	word := trim(input)
	escaped := escape(word)
	return len(escaped), len(word)
}

func trim(input string) string {
	return strings.Trim(input, " \t\r\n")
}

func escape(input string) string {
	return strconv.Quote(input)
}
