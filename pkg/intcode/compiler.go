package intcode

import "fmt"

// NewCompiler returns a new compiler.
func NewCompiler() *Compiler {
	return &Compiler{}
}

// Word is a single chunk of program memory.
type Word interface{}

// Literal is a value that's read directly.
type Literal int

// Symbol is a reference to a variable.
type Symbol int

// PC is the current length of the program, i.e. the PC value.
// The value itself is a relative position to the PC, i.e. forward N words or backwards N words.
// Use 0 to indicate the PC itself.
type PC int

// Parameter is an input to an operation.
type Parameter struct {
	Mode   int
	Value  int
	Symbol bool
	PC     bool
}

// Word returns the parameter as a word.
func (p Parameter) Word() Word {
	if p.Symbol {
		return Symbol(p.Value)
	}
	if p.PC {
		return PC(p.Value)
	}
	return Literal(p.Value)
}

// Compiler is a helper for writing intcode programs.
type Compiler struct {
	Program     []Word
	SymbolAddrs map[string]int
	Symbols     []int
}

// Compile processes the program, setting up variable references where relevant.
func (c *Compiler) Compile() (program []int) {
	for _, word := range c.Program {
		switch typed := word.(type) {
		case Literal:
			program = append(program, int(typed))
			continue
		case Symbol:
			program = append(program, len(c.Program)+int(typed)+1) // the symbol offset + the 1 for the halt.
			continue
		case PC:
			program = append(program, len(program)+int(typed))
		}
	}
	program = append(program, OpHalt)
	program = append(program, c.Symbols...)
	return
}

// CreateSymbol creates a symbol.
func (c *Compiler) CreateSymbol(name string) {
	if c.SymbolAddrs == nil {
		c.SymbolAddrs = make(map[string]int)
	}
	if _, ok := c.SymbolAddrs[name]; ok {
		panic(fmt.Sprintf("symbol already exists: %q", name))
	}
	c.SymbolAddrs[name] = len(c.Symbols)
	c.Symbols = append(c.Symbols, 0)
}

// SymbolValueParameter returns a new symbol parameter that is a immediate mode parameter.
// This will load the value at a given (named) address and use it for an operation.
func (c *Compiler) SymbolValueParameter(name string) Parameter {
	return Parameter{
		Symbol: true,
		Mode:   ParameterModeValue,
		Value:  c.SymbolAddrs[name],
	}
}

// SymbolReferenceParameter returns a new symbol parameter that is a reference mode parameter.
// This will load the value at a given (named) address and use that as the address to fetch the value to be used in the operation.
func (c *Compiler) SymbolReferenceParameter(name string) Parameter {
	return Parameter{
		Symbol: true,
		Mode:   ParameterModeReference,
		Value:  c.SymbolAddrs[name],
	}
}

// SymbolStoreParameter is a paramter to be used as the X register value
// that indicates it is a symbol address.
func (c *Compiler) SymbolStoreParameter(name string) Parameter {
	return Parameter{
		Symbol: true,
		Value:  c.SymbolAddrs[name],
	}
}

// ValueParameter returns an immediate mode parameter.
func (c *Compiler) ValueParameter(value int) Parameter {
	return Parameter{
		Mode:  ParameterModeValue,
		Value: value,
	}
}

// ReferenceParameter returns an reference mode parameter.
func (c *Compiler) ReferenceParameter(addr int) Parameter {
	return Parameter{
		Mode:  ParameterModeReference,
		Value: addr,
	}
}

// StoreParameter returns a program space storage parameter.
func (c *Compiler) StoreParameter(x int) Parameter {
	return Parameter{
		Value: x,
	}
}

// SetSymbolValue sets a symbol value, that is, it writes
// directly to the symbol address for a given name, a given value.
func (c *Compiler) SetSymbolValue(name string, value int) {
	c.Symbols[c.SymbolAddrs[name]] = value
}

// PC returns a PC parameter with a given offset.
func (c *Compiler) PC(offset int) Parameter {
	return Parameter{
		PC:    true,
		Value: offset,
	}
}

// Halt writes a halt operation at the current pc.
func (c *Compiler) Halt() {
	c.Program = append(c.Program, Literal(OpHalt))
}

// Add writes a new add operation with a given set of parameters.
func (c *Compiler) Add(a, b, x Parameter) {
	opcode := OpCode{Op: OpAdd, Modes: [3]int{0, b.Mode, a.Mode}}
	c.Program = append(c.Program, Literal(FormatOpCode(opcode)))
	c.Program = append(c.Program, a.Word())
	c.Program = append(c.Program, b.Word())
	c.Program = append(c.Program, x.Word())
}

// Mul writes a new mul operation with a given set of parameters.
func (c *Compiler) Mul(a, b, x Parameter) {
	opcode := OpCode{Op: OpAdd, Modes: [3]int{0, b.Mode, a.Mode}}

	c.Program = append(c.Program, Literal(FormatOpCode(opcode)))
	c.Program = append(c.Program, a.Word())
	c.Program = append(c.Program, b.Word())
	c.Program = append(c.Program, x.Word())
}

// Input writes an input with a given storage address.
func (c *Compiler) Input(x Parameter) {
	c.Program = append(c.Program, Literal(OpInput))
	c.Program = append(c.Program, x.Word())
}

// Print writes an print with a given parameter.
func (c *Compiler) Print(x Parameter) {
	opcode := OpCode{Op: OpPrint, Modes: [3]int{0, 0, x.Mode}}
	c.Program = append(c.Program, Literal(FormatOpCode(opcode)))
	c.Program = append(c.Program, x.Word())
}

// JumpIfTrue writes a new jump-if-true operation with a given set of parameters.
func (c *Compiler) JumpIfTrue(a, b Parameter) {
	opcode := OpCode{Op: OpJumpIfTrue, Modes: [3]int{0, b.Mode, a.Mode}}
	c.Program = append(c.Program, Literal(FormatOpCode(opcode)))
	c.Program = append(c.Program, a.Value)
	c.Program = append(c.Program, b.Value)
}

// JumpIfFalse writes a new jump-if-false operation with a given set of parameters.
func (c *Compiler) JumpIfFalse(a, b Parameter) {
	opcode := OpCode{Op: OpJumpIfFalse, Modes: [3]int{0, b.Mode, a.Mode}}
	c.Program = append(c.Program, Literal(FormatOpCode(opcode)))
	c.Program = append(c.Program, a.Word())
	c.Program = append(c.Program, b.Word())
}

// LessThan writes a new less-than operation with a given set of parameters.
func (c *Compiler) LessThan(a, b, x Parameter) {
	opcode := OpCode{Op: OpLessThan, Modes: [3]int{0, b.Mode, a.Mode}}
	c.Program = append(c.Program, Literal(FormatOpCode(opcode)))
	c.Program = append(c.Program, a.Word())
	c.Program = append(c.Program, b.Word())
	c.Program = append(c.Program, x.Word())
}

// Equals writes a new less-than operation with a given set of parameters.
func (c *Compiler) Equals(a, b, x Parameter) {
	opcode := OpCode{Op: OpEquals, Modes: [3]int{0, b.Mode, a.Mode}}
	c.Program = append(c.Program, Literal(FormatOpCode(opcode)))
	c.Program = append(c.Program, a.Word())
	c.Program = append(c.Program, b.Word())
	c.Program = append(c.Program, x.Word())
}
