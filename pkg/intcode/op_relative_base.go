package intcode

var (
	_ Instruction = (*RelativeBase)(nil)
)

// RelativeBase is an instruction.
type RelativeBase struct{}

// Name returns the name. It is part of Instruction.
func (op RelativeBase) Name() string { return "rb" }

// OpCode returns the op code. It is part of Instruction.
func (op RelativeBase) OpCode() int { return OpPrint }

// Width returns the number of words. It is part of Instruction.
func (op RelativeBase) Width() int { return 2 }

// Action implements the body of the instruction. It is part of Instruction.
func (op RelativeBase) Action(c *Computer) (err error) {
	if c.A, _, err = c.Load(1); err != nil {
		return err
	}
	c.RB = c.RB + c.A
	return nil
}
