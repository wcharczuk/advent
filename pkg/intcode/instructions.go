package intcode

// Instructions is the full list of instructions.
var (
	Instructions = map[int64]Instruction{
		OpHalt:         Halt{},
		OpAdd:          Add{},
		OpMul:          Mul{},
		OpInput:        Input{},
		OpPrint:        Print{},
		OpJumpIfTrue:   JumpIfTrue{},
		OpJumpIfFalse:  JumpIfFalse{},
		OpLessThan:     LessThan{},
		OpEquals:       Equals{},
		OpRelativeBase: RelativeBase{},
	}
)

// InstructionNames is the full list of instructions organized by name.
var (
	InstructionNames = map[string]Instruction{}
)
