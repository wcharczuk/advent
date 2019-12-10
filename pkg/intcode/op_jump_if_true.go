package intcode

var (
	_ Instruction   = (*JumpIfTrue)(nil)
	_ InstructionPC = (*JumpIfTrue)(nil)
)

// JumpIfTrue is an instruction.
type JumpIfTrue struct{}

// Name returns the name. It is part of Instruction.
func (op JumpIfTrue) Name() string { return "jump-if-true" }

// OpCode returns the op code. It is part of Instruction.
func (op JumpIfTrue) OpCode() int { return OpJumpIfTrue }

// Width returns the number of words. It is part of Instruction.
func (op JumpIfTrue) Width() int { return 3 }

// Action implements the body of the instruction. It is part of Instruction.
func (op JumpIfTrue) Action(c *Computer) (err error) {
	if c.A, _, err = c.Load(1); err != nil {
		return err
	}
	if c.B, _, err = c.Load(2); err != nil {
		return err
	}
	return nil
}

// MovePC implements InstructionPC
func (op JumpIfTrue) MovePC(c *Computer) error {
	if c.A > 0 {
		c.PC = c.B
		return nil
	}
	c.PC = c.PC + int64(op.Width())
	return nil
}
