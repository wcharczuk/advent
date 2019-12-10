package intcode

var (
	_ Instruction = (*LessThan)(nil)
)

// LessThan is an instruction.
type LessThan struct{}

// Name returns the name. It is part of Instruction.
func (op LessThan) Name() string { return "less-than" }

// OpCode returns the op code. It is part of Instruction.
func (op LessThan) OpCode() int { return OpLessThan }

// Width returns the number of words. It is part of Instruction.
func (op LessThan) Width() int { return 4 }

// Action implements the body of the instruction. It is part of Instruction.
func (op LessThan) Action(c *Computer) (err error) {
	if c.A, _, err = c.Load(1); err != nil {
		return err
	}
	if c.B, _, err = c.Load(2); err != nil {
		return err
	}
	if c.X, _, err = c.LoadForStore(3); err != nil {
		return err
	}
	if c.A < c.B {
		c.Store(1)
	} else {
		c.Store(0)
	}
	return nil
}
