package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	codeFile := "../testdata/day12"

	if f, err := os.Open(codeFile); err == nil {
		defer f.Close()

		contents, readErr := ioutil.ReadAll(f)
		if readErr != nil {
			fmt.Errorf("Error reading code file: %v", readErr)
		}

		total := processFile(contents)
		fmt.Printf("Total in file: %d\n", total)
	} else {
		fmt.Errorf("Error opening code file: %v", err)
		os.Exit(1)
	}
}

func processFile(fileContents []byte) int {
	total := 0
	state := 0

	number := ""
	for x := 0; x < len(fileContents); x++ {
		c := string(fileContents[x])
		switch state {
		case 0:
			if c == `"` {
				state = 1
			} else if isNumberCharacter(c) {
				number = number + c
				state = 2
			}
		case 1:
			if c == `"` {
				state = 0
			}
		case 2:
			if !isNumberCharacter(c) {
				value, _ := strconv.Atoi(number)
				total += value
				state = 0
				number = ""
			} else {
				number = number + c
			}
		}
	}
	return total
}

func isNumberCharacter(input string) bool {
	switch input {
	case "-", ".", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
		return true
	default:
		return false
	}
}
