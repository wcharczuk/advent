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
	Code   int
	Memory int
}

func main() {
	dataFile := "../testdata/day8"

	testCases := []testCase{
		{`""`, 2, 0},
		{`"abc"`, 5, 3},
		{`"aaa\"aaa"`, 10, 7},
		{`"\x27"`, 6, 1},
		{`"\"lhcaurdqzyjyu"`, 17, 14},
		{`"inleep\\mgl"`, 13, 10},
	}

	allPassed := true
	for _, test := range testCases {
		code, memory := process(test.Value)
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
			code, memory := process(word)
			codeCount += code
			memoryCount += memory
		}
	}

	fmt.Println("Code:", codeCount, "Memory:", memoryCount, "Diff:", codeCount-memoryCount)
}

func trim(input string) string {
	return strings.Trim(input, " \t\r\n")
}

func process(input string) (int, int) {
	word := trim(input)
	escaped := escape(removeOuterQuotes(word))
	return len(word), len(escaped)
}

func removeOuterQuotes(input string) string {
	return input[1 : len(input)-1]
}

func escape(input string) string {
	output := ""

	hexValue := ""
	state := 0

	for x := 0; x < len(input); x++ {
		c := string(input[x])
		switch state {
		case 0:
			if c == "\\" {
				state = 1
			} else {
				output = output + c
			}
		case 1:
			if c == "\"" {
				output = output + c
				state = 0
			} else if c == "\\" {
				output = output + c
				state = 0
			} else if c == "x" {
				hexValue = "\\x"
				state = 2
			} else {
				output = output + "\\"
				output = output + c
				state = 0
			}
			break
		case 2:
			hexValue = hexValue + c
			state = 3
			break
		case 3:
			hexValue = hexValue + c
			escaped, _ := strconv.Unquote(fmt.Sprintf(`"%s"`, hexValue))
			output = output + escaped
			state = 0
		}
	}

	return output
}
