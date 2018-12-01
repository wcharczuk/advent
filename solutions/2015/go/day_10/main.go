package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if process("1", 1) != "11" {
		fmt.Println(process("1", 1), "!=", "11")
		os.Exit(1)
	}

	if process("11", 1) != "21" {
		fmt.Println(process("11", 1), "!=", "21")
		os.Exit(1)
	}

	if process("21", 1) != "1211" {
		fmt.Println(process("21", 1), "!=", "1211")
		os.Exit(1)
	}

	if process("1211", 1) != "111221" {
		fmt.Println(process("1211", 1), "!=", "111221")
		os.Exit(1)
	}

	if process("111221", 1) != "312211" {
		fmt.Println(process("111221", 1), "!=", "312211")
		os.Exit(1)
	}

	if process("1", 5) != "312211" {
		fmt.Println(process("1", 5), "!=", "312211")
		os.Exit(1)
	}

	final := process("1113222113", 40)
	println("final value length", len(final))
}

func process(input string, iterations int) string {
	working := make([]byte, len(input))
	copy(working, []byte(input))

	workingOutput := []byte{}
	for x := 0; x < iterations; x++ {
		for y := 0; y < len(working); {
			r := y
			for r < len(working) && working[y] == working[r] {
				r++
			}
			workingOutput = append(workingOutput, byte(strconv.Itoa(r - y)[0]))
			workingOutput = append(workingOutput, working[y])
			y = r
		}
		working = workingOutput
		workingOutput = []byte{}
	}

	return string(working)
}
