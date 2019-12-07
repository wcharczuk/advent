package intcode

import (
	"fmt"
	"strconv"
	"strings"
)

// ParseOpCode parses an input as a structured opcode.
// The op code can be in the form
// <1 op>
// <10 op><1 op>
// <a mode><10 op><1 op>
// <b mode><a mode><20 op>
// <x mode><b mode><a mode><10 op><1 op>
func ParseOpCode(value int) (output OpCode, err error) {
	runes := []rune(strconv.Itoa(value))
	if len(runes) == 1 {
		runes = append([]rune(string("0000")), runes...)
	} else if len(runes) == 2 {
		runes = append([]rune(string("000")), runes...)
	} else if len(runes) == 3 {
		runes = append([]rune(string("00")), runes...)
	} else if len(runes) == 4 {
		runes = append([]rune(string("0")), runes...)
	}

	output.Op, err = strconv.Atoi(string([]rune{runes[3], runes[4]}))
	if err != nil {
		return
	}

	output.Modes = [3]int{0, 0, 0}
	if runes[0] == '1' {
		output.Modes[0] = 1
	}
	if runes[1] == '1' {
		output.Modes[1] = 1
	}
	if runes[2] == '1' {
		output.Modes[2] = 1
	}

	return
}

// FormatOpCode formats an op code as an integer.
func FormatOpCode(oc OpCode) int {
	if oc.Modes[0] == 0 && oc.Modes[1] == 0 && oc.Modes[2] == 0 {
		return oc.Op
	}

	var pieces []string
	if oc.Modes[2] != 0 {
		pieces = append(pieces, strconv.Itoa(oc.Modes[2]))
	}

	if oc.Modes[1] != 0 {
		pieces = append(pieces, strconv.Itoa(oc.Modes[1]))
	} else if len(pieces) > 0 {
		pieces = append(pieces, strconv.Itoa(oc.Modes[1]))
	}

	if oc.Modes[0] != 0 {
		pieces = append(pieces, strconv.Itoa(oc.Modes[0]))
	} else if len(pieces) > 0 {
		pieces = append(pieces, strconv.Itoa(oc.Modes[0]))
	}

	pieces = append(pieces, fmt.Sprintf("%02d", oc.Op))
	value, _ := strconv.Atoi(strings.Join(pieces, ""))
	return value
}

// OpCode is an operation.
type OpCode struct {
	Op    int
	Modes [3]int
}

// String returns a descriptive string for the op code.
func (oc OpCode) String() string {
	switch oc.Op {
	case OpHalt:
		return "halt"
	case OpAdd:
		return "add"
	case OpMul:
		return "mul"
	case OpInput:
		return "input"
	case OpPrint:
		return "print"
	case OpJumpIfTrue:
		return "jump if true"
	case OpJumpIfFalse:
		return "jump if false"
	case OpLessThan:
		return "less than"
	case OpEquals:
		return "equals"
	default:
		return "unknown"
	}
}

// Mode returns the mode for an index.
func (oc OpCode) Mode(index int) int {
	switch index {
	case 0:
		return oc.Modes[2]
	case 1:
		return oc.Modes[1]
	case 2:
		return oc.Modes[0]
	default:
		return 0
	}
}
