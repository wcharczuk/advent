package intcode

var (
	_ Instruction = (*Halt)(nil)
)

// Halt is an instruction.
type Halt struct{}

// Name returns the name. It is part of Instruction.
func (op Halt) Name() string { return "halt" }

// OpCode returns the op code. It is part of Instruction.
func (op Halt) OpCode() int { return OpHalt }

// Width returns the number of words. It is part of Instruction.
func (op Halt) Width() int { return 1 }

// Action implements the body of the instruction. It is part of Instruction.
func (op Halt) Action(c *Computer) (err error) {
	return ErrHalt
}
