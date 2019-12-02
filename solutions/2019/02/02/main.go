package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	contents, err := ioutil.ReadFile("../input")
	if err != nil {
		log.Fatal(err)
	}

	rawValues := strings.Split(string(contents), ",")
	opCodes := make([]int, len(rawValues))

	for x := 0; x < len(rawValues); x++ {
		opCodes[x], err = strconv.Atoi(strings.TrimSpace(rawValues[x]))
		if err != nil {
			log.Fatal(err)
		}
	}

	for noun := 0; noun < 99; noun++ {
		for verb := 0; verb < 99; verb++ {
			// make a copy
			testCodes := make([]int, len(opCodes))
			copy(testCodes, opCodes)

			testCodes[1] = noun
			testCodes[2] = verb

			answer, err := Intcode(testCodes...)
			if err != nil {
				log.Fatal(err)
			}
			if answer[0] == 19690720 {
				fmt.Printf("ANSWER: %d", (100*noun)+verb)
				os.Exit(0)
			}
		}
	}
	fmt.Fprintf(os.Stderr, "exhausted search space")
	os.Exit(1)
}

// Intcode processes the input.
func Intcode(opCodeInput ...int) (opCodes []int, err error) {
	opCodes = make([]int, len(opCodeInput))
	copy(opCodes, opCodeInput)

	var opCode, aValue, bValue, storeAt int
	for head := 0; head < len(opCodes); {
		opCode = opCodes[head]
		switch opCode {
		case 99:
			return
		case 1:
			// do add
			aValue = opCodes[opCodes[head+1]]
			bValue = opCodes[opCodes[head+2]]
			storeAt = opCodes[head+3]

			opCodes[storeAt] = aValue + bValue
			head = head + 4
		case 2:
			// do add
			aValue = opCodes[opCodes[head+1]]
			bValue = opCodes[opCodes[head+2]]
			storeAt = opCodes[head+3]

			opCodes[storeAt] = aValue * bValue
			head = head + 4
		default:
			err = fmt.Errorf("invalid opcode: %d", opCode)
			return
		}
	}
	return
}
