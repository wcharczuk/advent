package main

import (
	"fmt"
	"io/ioutil"
	"log"
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

	_, err = Intcode(opCodes...)
	if err != nil {
		log.Fatal(err)
	}
}

// Intcode processes the input.
func Intcode(opCodeInput ...int) (opCodes []int, err error) {
	opCodes = make([]int, len(opCodeInput))
	copy(opCodes, opCodeInput)

	var opCode OpCode
	var aValue, bValue, storeAt int

	for head := 0; head < len(opCodes); {
		opCode = ParseOpcode(opCodes[head])
		switch opCode.Op {
		case 99:
			return
		case 1: // add
			// do add
			aValue = opCode.Load(opCodes, 0, head+1)
			bValue = opCode.Load(opCodes, 1, head+2)
			storeAt = opCodes[head+3]

			opCodes[storeAt] = aValue + bValue
			head = head + 4
		case 2: // multiply
			// do add
			aValue = opCode.Load(opCodes, 0, head+1)
			bValue = opCode.Load(opCodes, 1, head+2)
			storeAt = opCodes[head+3]

			opCodes[storeAt] = aValue * bValue
			head = head + 4
		case 3: // input
			storeAt = opCodes[head+1]
			opCodes[storeAt] = readInt()
			head = head + 2
		case 4: // print
			storeAt = opCode.Load(opCodes, 0, head+1)
			fmt.Println(storeAt)
			head = head + 2
		case 5: // jump if true
			aValue = opCode.Load(opCodes, 0, head+1)
			bValue = opCode.Load(opCodes, 1, head+2)
			if aValue > 0 {
				head = bValue
				continue
			}
			head = head + 3
		case 6: // jump if false
			aValue = opCode.Load(opCodes, 0, head+1)
			bValue = opCode.Load(opCodes, 1, head+2)
			if aValue == 0 {
				head = bValue
				continue
			}
			head = head + 3
		case 7: // less than
			aValue = opCode.Load(opCodes, 0, head+1)
			bValue = opCode.Load(opCodes, 1, head+2)
			storeAt = opCodes[head+3]
			if aValue < bValue {
				opCodes[storeAt] = 1
			} else {
				opCodes[storeAt] = 0
			}
			head = head + 4
		case 8: // equals
			aValue = opCode.Load(opCodes, 0, head+1)
			bValue = opCode.Load(opCodes, 1, head+2)
			storeAt = opCodes[head+3]
			if aValue == bValue {
				opCodes[storeAt] = 1
			} else {
				opCodes[storeAt] = 0
			}
			head = head + 4
		default:
			err = fmt.Errorf("invalid opcode: %d", opCode)
			return
		}
	}
	return
}

// ParseOpcode parses an input as a structured opcode.
func ParseOpcode(opCodeInput int) OpCode {
	runes := []rune(strconv.Itoa(opCodeInput))
	if len(runes) < 3 {
		return OpCode{Op: opCodeInput, Modes: make([]int, 3)}
	}
	rl := len(runes)
	rle2 := rl - 2
	rle1 := rl - 1
	opCode, _ := strconv.Atoi(string([]rune{runes[rle2], runes[rle1]}))
	modes := make([]int, 3)

	if len(runes) > 2 {
		for x := 0; x < len(runes)-2; x++ {
			if runes[rl-(x+3)] == '1' {
				modes[x] = 1
			} else {
				modes[x] = 0
			}
		}
	}

	return OpCode{
		Op:    opCode,
		Modes: modes,
	}
}

// OpCode is an operation.
type OpCode struct {
	Op    int
	Modes []int
}

// Mode returns the mode for an index.
func (oc OpCode) Mode(index int) int {
	return oc.Modes[index]
}

// Load loads a value with a given param index, at a given addr.
func (oc OpCode) Load(opCodes []int, index, addr int) int {
	switch oc.Mode(index) {
	case 0:
		return opCodes[opCodes[addr]]
	case 1:
		return opCodes[addr]
	default:
		panic("invalid mode")
	}
}

func readInt() int {
	fmt.Print("Input: ")
	var i int
	fmt.Scanf("%d", &i)
	return i
}
