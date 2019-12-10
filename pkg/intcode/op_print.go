package intcode

import "fmt"

var (
	_ Instruction = (*Print)(nil)
)

// Print is an instruction.
type Print struct{}

// Name returns the name. It is part of Instruction.
func (op Print) Name() string { return "print" }

// OpCode returns the op code. It is part of Instruction.
func (op Print) OpCode() int { return OpPrint }

// Width returns the number of words. It is part of Instruction.
func (op Print) Width() int { return 2 }

// Action implements the body of the instruction. It is part of Instruction.
func (op Print) Action(c *Computer) (err error) {
	if c.A, _, err = c.Load(1); err != nil {
		return err
	}
	if len(c.OutputHandlers) > 0 {
		for _, handler := range c.OutputHandlers {
			handler(c.A)
		}
	} else {
		fmt.Println(c.A)
	}
	return nil
}
