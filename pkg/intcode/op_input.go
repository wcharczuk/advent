package intcode

var (
	_ Instruction = (*Input)(nil)
)

// Input is an instruction.
type Input struct{}

// Name returns the name. It is part of Instruction.
func (op Input) Name() string { return "input" }

// OpCode returns the op code. It is part of Instruction.
func (op Input) OpCode() int { return OpInput }

// Width returns the number of words. It is part of Instruction.
func (op Input) Width() int { return 2 }

// Action implements the body of the instruction. It is part of Instruction.
func (op Input) Action(c *Computer) (err error) {
	if c.X, _, err = c.LoadForStore(1); err != nil {
		return err
	}
	var value int64
	if c.InputHandler != nil {
		value = c.InputHandler()
	} else {
		value = ReadInt()
	}
	if err = c.Store(value); err != nil {
		return err
	}
	return nil
}
