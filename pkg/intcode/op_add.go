package intcode

var (
	_ Instruction = (*Add)(nil)
)

// Add is an instruction.
type Add struct{}

// Name returns the name. It is part of Instruction.
func (op Add) Name() string { return "add" }

// OpCode returns the op code. It is part of Instruction.
func (op Add) OpCode() int { return OpAdd }

// Width returns the number of words. It is part of Instruction.
func (op Add) Width() int { return 4 }

// Action implements the body of the instruction. It is part of Instruction.
func (op Add) Action(c *Computer) (err error) {
	if c.A, _, err = c.Load(1); err != nil {
		return err
	}
	if c.B, _, err = c.Load(2); err != nil {
		return err
	}
	if c.X, _, err = c.LoadForStore(3); err != nil {
		return err
	}
	if err = c.Store(c.A + c.B); err != nil {
		return err
	}
	return nil
}
