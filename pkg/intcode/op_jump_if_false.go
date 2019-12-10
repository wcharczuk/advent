package intcode

var (
	_ Instruction   = (*JumpIfFalse)(nil)
	_ InstructionPC = (*JumpIfFalse)(nil)
)

// JumpIfFalse is an instruction.
type JumpIfFalse struct{}

// Name returns the name. It is part of Instruction.
func (op JumpIfFalse) Name() string { return "jump-if-false" }

// OpCode returns the op code. It is part of Instruction.
func (op JumpIfFalse) OpCode() int { return OpJumpIfFalse }

// Width returns the number of words. It is part of Instruction.
func (op JumpIfFalse) Width() int { return 3 }

// Action implements the body of the instruction. It is part of Instruction.
func (op JumpIfFalse) Action(c *Computer) (err error) {
	if c.A, _, err = c.Load(1); err != nil {
		return err
	}
	if c.B, _, err = c.Load(2); err != nil {
		return err
	}
	return nil
}

// MovePC implements InstructionPC
func (op JumpIfFalse) MovePC(c *Computer) error {
	if c.A == 0 {
		c.PC = c.B
		return nil
	}
	c.PC = c.PC + int64(op.Width())
	return nil
}
