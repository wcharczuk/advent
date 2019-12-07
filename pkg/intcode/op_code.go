package intcode

import "strconv"

// ParseOpCode parses an input as a structured opcode.
func ParseOpCode(opCodeInput int) OpCode {
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
		return "less"
	case OpEquals:
		return "equal"
	default:
		return "unknown"
	}
}

// Mode returns the mode for an index.
func (oc OpCode) Mode(index int) int {
	return oc.Modes[index]
}
