package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/wcharczuk/advent/pkg/dbg"
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

	// weird subs?
	opCodes[1] = 12
	opCodes[2] = 2

	answer, err := Intcode(opCodes...)
	if err != nil {
		log.Fatal(err)
	}

	dbg.Spew(answer)
	fmt.Printf("ANSWER: %d", answer[0])
}

// Intcode processes the input.
func Intcode(opCodeInput ...int) (opCodes []int, err error) {
	// make a copy
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
