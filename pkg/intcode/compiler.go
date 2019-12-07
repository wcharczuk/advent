package intcode

// Value returns a value parameter.
func Value(value int) Parameter {
	return Parameter{Mode: ParameterModeReference, Value: value}
}

// Reference returns a reference parameter.
func Reference(addr int) Parameter {
	return Parameter{Mode: ParameterModeReference, Value: addr}
}

// Parameter is used to tell the computer to take
// a value or a reference based on the mode.
type Parameter struct {
	Mode  int
	Value int
}

// NewCompiler returns a new compiler.
func NewCompiler(memorySize int) *Compiler {
	return &Compiler{}
}

// Compiler is a helper for writing intcode programs.
type Compiler struct {
	// Program is the program as written.
	Program []int
}

// Compile processes the program, setting up variable references where relevant.
func (c *Compiler) Compile() ([]int, error) {
	return c.Program, nil
}

// Halt writes a halt operation at the current pc.
func (c *Compiler) Halt() {
	c.Program = append(c.Program, OpHalt)
}

// Add writes a new add operation with a given set of parameters.
func (c *Compiler) Add(a, b Parameter, x int) {
	opcode := OpCode{Op: OpAdd, Modes: [3]int{0, b.Mode, a.Mode}}
	c.Program = append(c.Program, FormatOpCode(opcode))
	c.Program = append(c.Program, a.Value)
	c.Program = append(c.Program, b.Value)
	c.Program = append(c.Program, x)
}

// Mul writes a new mul operation with a given set of parameters.
func (c *Compiler) Mul(a, b Parameter, x int) {
	opcode := OpCode{Op: OpAdd, Modes: [3]int{0, b.Mode, a.Mode}}
	c.Program = append(c.Program, FormatOpCode(opcode))
	c.Program = append(c.Program, a.Value)
	c.Program = append(c.Program, b.Value)
	c.Program = append(c.Program, x)
}

// Input writes an input with a given storage address.
func (c *Compiler) Input(x int) {
	c.Program = append(c.Program, OpInput)
	c.Program = append(c.Program, x)
}

// Print writes an print with a given parameter.
func (c *Compiler) Print(x Parameter) {
	opcode := OpCode{Op: OpPrint, Modes: [3]int{0, 0, x.Mode}}
	c.Program = append(c.Program, FormatOpCode(opcode))
	c.Program = append(c.Program, x.Value)
}

// JumpIfTrue writes a new jump-if-true operation with a given set of parameters.
func (c *Compiler) JumpIfTrue(a, b Parameter) {
	opcode := OpCode{Op: OpJumpIfTrue, Modes: [3]int{0, b.Mode, a.Mode}}
	c.Program = append(c.Program, FormatOpCode(opcode))
	c.Program = append(c.Program, a.Value)
	c.Program = append(c.Program, b.Value)
}

// JumpIfFalse writes a new jump-if-false operation with a given set of parameters.
func (c *Compiler) JumpIfFalse(a, b Parameter) {
	opcode := OpCode{Op: OpJumpIfFalse, Modes: [3]int{0, b.Mode, a.Mode}}
	c.Program = append(c.Program, FormatOpCode(opcode))
	c.Program = append(c.Program, a.Value)
	c.Program = append(c.Program, b.Value)
}

// LessThan writes a new less-than operation with a given set of parameters.
func (c *Compiler) LessThan(a, b Parameter, x int) {
	opcode := OpCode{Op: OpLessThan, Modes: [3]int{0, b.Mode, a.Mode}}
	c.Program = append(c.Program, FormatOpCode(opcode))
	c.Program = append(c.Program, a.Value)
	c.Program = append(c.Program, b.Value)
	c.Program = append(c.Program, x)
}

// Equals writes a new less-than operation with a given set of parameters.
func (c *Compiler) Equals(a, b Parameter, x int) {
	opcode := OpCode{Op: OpEquals, Modes: [3]int{0, b.Mode, a.Mode}}
	c.Program = append(c.Program, FormatOpCode(opcode))
	c.Program = append(c.Program, a.Value)
	c.Program = append(c.Program, b.Value)
	c.Program = append(c.Program, x)
}
