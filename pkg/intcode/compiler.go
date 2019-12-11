package intcode

// NewCompiler returns a new compiler.
func NewCompiler() *Compiler {
	return &Compiler{}
}

// Parameter is an input to an operation.
type Parameter struct {
	Mode  int
	Value int64
}

// Compiler is a helper for writing intcode programs.
type Compiler struct {
	Program []int64
}

// Compile processes the program, setting up variable references where relevant.
func (c *Compiler) Compile() (program []int64) {
	for _, word := range c.Program {
		program = append(program, int64(word))
	}
	return
}

// Value returns an immediate mode parameter.
func (c *Compiler) Value(value int64) Parameter {
	return Parameter{
		Mode:  ParameterModeValue,
		Value: value,
	}
}

// Reference returns an reference mode parameter.
func (c *Compiler) Reference(addr int64) Parameter {
	return Parameter{
		Mode:  ParameterModeReference,
		Value: addr,
	}
}

// Relative returns an relative mode parameter.
func (c *Compiler) Relative(offset int64) Parameter {
	return Parameter{
		Mode:  ParameterModeRelative,
		Value: offset,
	}
}

// EmitHalt writes a halt operation at the current pc.
func (c *Compiler) EmitHalt() {
	c.Program = append(c.Program, OpHalt)
}

// EmitAdd writes a new add operation with a given set of parameters.
func (c *Compiler) EmitAdd(a, b, x Parameter) {
	opcode := OpCode{Op: OpAdd, Modes: [3]int{x.Mode, b.Mode, a.Mode}}
	c.Program = append(c.Program, FormatOpCode(opcode))
	c.Program = append(c.Program, a.Value)
	c.Program = append(c.Program, b.Value)
	c.Program = append(c.Program, x.Value)
}

// EmitMul writes a new mul operation with a given set of parameters.
func (c *Compiler) EmitMul(a, b, x Parameter) {
	opcode := OpCode{Op: OpAdd, Modes: [3]int{x.Mode, b.Mode, a.Mode}}
	c.Program = append(c.Program, FormatOpCode(opcode))
	c.Program = append(c.Program, a.Value)
	c.Program = append(c.Program, b.Value)
	c.Program = append(c.Program, x.Value)
}

// EmitInput writes an input with a given storage address.
func (c *Compiler) EmitInput(x Parameter) {
	opcode := OpCode{Op: OpAdd, Modes: [3]int{0, 0, x.Mode}}
	c.Program = append(c.Program, FormatOpCode(opcode))
	c.Program = append(c.Program, x.Value)
}

// EmitPrint writes an print with a given parameter.
func (c *Compiler) EmitPrint(x Parameter) {
	opcode := OpCode{Op: OpPrint, Modes: [3]int{0, 0, x.Mode}}
	c.Program = append(c.Program, FormatOpCode(opcode))
	c.Program = append(c.Program, x.Value)
}

// EmitJumpIfTrue writes a new jump-if-true operation with a given set of parameters.
func (c *Compiler) EmitJumpIfTrue(testValue, jumpTo Parameter) {
	opcode := OpCode{Op: OpJumpIfTrue, Modes: [3]int{0, jumpTo.Mode, testValue.Mode}}
	c.Program = append(c.Program, FormatOpCode(opcode))
	c.Program = append(c.Program, testValue.Value)
	c.Program = append(c.Program, jumpTo.Value)
}

// EmitJumpIfFalse writes a new jump-if-false operation with a given set of parameters.
func (c *Compiler) EmitJumpIfFalse(testValue, jumpTo Parameter) {
	opcode := OpCode{Op: OpJumpIfFalse, Modes: [3]int{0, jumpTo.Mode, testValue.Mode}}
	c.Program = append(c.Program, FormatOpCode(opcode))
	c.Program = append(c.Program, testValue.Value)
	c.Program = append(c.Program, jumpTo.Value)
}

// EmitLessThan writes a new less-than operation with a given set of parameters.
func (c *Compiler) EmitLessThan(a, b, x Parameter) {
	opcode := OpCode{Op: OpLessThan, Modes: [3]int{x.Mode, b.Mode, a.Mode}}
	c.Program = append(c.Program, FormatOpCode(opcode))
	c.Program = append(c.Program, a.Value)
	c.Program = append(c.Program, b.Value)
	c.Program = append(c.Program, x.Value)
}

// EmitEquals writes a new less-than operation with a given set of parameters.
func (c *Compiler) EmitEquals(a, b, x Parameter) {
	opcode := OpCode{Op: OpEquals, Modes: [3]int{x.Mode, b.Mode, a.Mode}}
	c.Program = append(c.Program, FormatOpCode(opcode))
	c.Program = append(c.Program, a.Value)
	c.Program = append(c.Program, b.Value)
	c.Program = append(c.Program, x.Value)
}

// HelperEmitSet is a helper that emits an `Add(0, A, X)`, effectively
// setting the value X to A,
func (c *Compiler) HelperEmitSet(a, x Parameter) {
	opcode := OpCode{Op: OpAdd, Modes: [3]int{x.Mode, a.Mode, ParameterModeValue}}
	c.Program = append(c.Program, FormatOpCode(opcode))
	c.Program = append(c.Program, 0)
	c.Program = append(c.Program, a.Value)
	c.Program = append(c.Program, x.Value)
}

// EmitRelativeBase writes an input with a given storage address.
func (c *Compiler) EmitRelativeBase(x Parameter) {
	opcode := OpCode{Op: OpRelativeBase, Modes: [3]int{0, 0, x.Mode}}
	c.Program = append(c.Program, FormatOpCode(opcode))
	c.Program = append(c.Program, x.Value)
}

// HelperEmitLiteral emits a literal to the program, returning the index of that literal.
// This typically used to dump values to the program directly that will
// be jumped around.
func (c *Compiler) HelperEmitLiteral(value int64) int {
	index := len(c.Program)
	c.Program = append(c.Program, value)
	return index
}
