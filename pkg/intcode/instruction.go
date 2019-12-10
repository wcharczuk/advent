package intcode

// Instruction is a decoded op instruction.
type Instruction interface {
	Name() string
	OpCode() int
	Width() int
	Action(c *Computer) error
}

// InstructionPC implements a PC management step.
// This is typically used in jump instructions.
type InstructionPC interface {
	MovePC(c *Computer) error
}
