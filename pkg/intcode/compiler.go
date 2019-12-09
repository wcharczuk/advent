package intcode

import "fmt"

// NewCompiler returns a new compiler.
func NewCompiler() *Compiler {
	return &Compiler{}
}

// Word is a single chunk of program memory.
type Word interface{}

// Literal is an offset from the beginning of the program, i.e. PC@0.
type Literal int64

// Symbol is an offset from the end of the program,
// typically it is used to give the value of a symbol (i.e. a variable)
// but can also be used to move relative to the end of the program.
// If your last instruction is 99 (HALT), you can use Symbol(-1) to jump to that halt.
type Symbol int64

// PC is an offset from the current PC register,
// You can use it to move relative to your progress within a program.
type PC int64

// Parameter is an input to an operation.
type Parameter struct {
	Mode   int
	Value  int64
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
	SymbolAddrs map[string]int64
	Symbols     []int64
}

// Compile processes the program, setting up variable references where relevant.
func (c *Compiler) Compile() (program []int64) {
	for _, word := range c.Program {
		switch typed := word.(type) {
		case Literal:
			program = append(program, int64(typed))
			continue
		case Symbol:
			program = append(program, int64(len(c.Program))+(int64(typed)+1))
			continue
		case PC:
			program = append(program, int64(len(program))+int64(typed))
			continue
		}
	}
	program = append(program, OpHalt)
	program = append(program, c.Symbols...)
	return
}

// CreateSymbol creates a symbol.
func (c *Compiler) CreateSymbol(name string) {
	if c.SymbolAddrs == nil {
		c.SymbolAddrs = make(map[string]int64)
	}
	if _, ok := c.SymbolAddrs[name]; ok {
		panic(fmt.Sprintf("symbol already exists: %q", name))
	}
	c.SymbolAddrs[name] = int64(len(c.Symbols))
	c.Symbols = append(c.Symbols, 0)
}

// ValueSymbol returns a new symbol parameter that is a immediate mode parameter.
// This will load the value at a given (named) address and use it for an operation.
func (c *Compiler) ValueSymbol(name string) Parameter {
	return Parameter{
		Symbol: true,
		Mode:   ParameterModeValue,
		Value:  c.SymbolAddrs[name],
	}
}

// ReferenceSymbol returns a new symbol parameter that is a reference mode parameter.
// This will load the value at a given (named) address and use that as the address to fetch the value to be used in the operation.
func (c *Compiler) ReferenceSymbol(name string) Parameter {
	return Parameter{
		Symbol: true,
		Mode:   ParameterModeReference,
		Value:  c.SymbolAddrs[name],
	}
}

// StoreSymbol is a paramter to be used as the X register value
// that indicates it is a symbol address.
func (c *Compiler) StoreSymbol(name string) Parameter {
	return Parameter{
		Symbol: true,
		Value:  c.SymbolAddrs[name],
	}
}

// ValueSymbolOffset returns a value that represents the address
// that is offset from the end of the program.
func (c *Compiler) ValueSymbolOffset(offset int64) Parameter {
	return Parameter{
		Mode:   ParameterModeValue,
		Symbol: true,
		Value:  offset,
	}
}

// ReferenceSymbolOffset returns a value that represents the address
// that is offset from the end of the program.
func (c *Compiler) ReferenceSymbolOffset(offset int64) Parameter {
	return Parameter{
		Mode:   ParameterModeReference,
		Symbol: true,
		Value:  offset,
	}
}

// ValueHaltAddr is a helper that returns the address of the final `99 (HALT)`
// instruction that is added by the compiler.
func (c *Compiler) ValueHaltAddr() Parameter {
	return Parameter{
		Mode:   ParameterModeValue,
		Symbol: true,
		Value:  -1,
	}
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

// Store returns a program space storage parameter.
// It is intended to be used as an address.
func (c *Compiler) Store(addr int64) Parameter {
	return Parameter{
		Value: addr,
	}
}

// ValuePC retruns a parameter that is a value of the current PC address with a given offset.
func (c *Compiler) ValuePC(offset int64) Parameter {
	return Parameter{
		Mode:  ParameterModeValue,
		PC:    true,
		Value: offset,
	}
}

// ReferencePC returns a PC parameter with a given offset.
// It is always a reference, that is, it will have the effect
// of returning the value at a given address, not the address itself.
func (c *Compiler) ReferencePC(offset int64) Parameter {
	return Parameter{
		Mode:  ParameterModeReference,
		PC:    true,
		Value: offset,
	}
}

// EmitHalt writes a halt operation at the current pc.
func (c *Compiler) EmitHalt() {
	c.Program = append(c.Program, Literal(OpHalt))
}

// EmitAdd writes a new add operation with a given set of parameters.
func (c *Compiler) EmitAdd(a, b, x Parameter) {
	opcode := OpCode{Op: OpAdd, Modes: [3]int{0, b.Mode, a.Mode}}
	c.Program = append(c.Program, Literal(FormatOpCode(opcode)))
	c.Program = append(c.Program, a.Word())
	c.Program = append(c.Program, b.Word())
	c.Program = append(c.Program, x.Word())
}

// EmitMul writes a new mul operation with a given set of parameters.
func (c *Compiler) EmitMul(a, b, x Parameter) {
	opcode := OpCode{Op: OpAdd, Modes: [3]int{0, b.Mode, a.Mode}}

	c.Program = append(c.Program, Literal(FormatOpCode(opcode)))
	c.Program = append(c.Program, a.Word())
	c.Program = append(c.Program, b.Word())
	c.Program = append(c.Program, x.Word())
}

// EmitInput writes an input with a given storage address.
func (c *Compiler) EmitInput(x Parameter) {
	c.Program = append(c.Program, Literal(OpInput))
	c.Program = append(c.Program, x.Word())
}

// EmitPrint writes an print with a given parameter.
func (c *Compiler) EmitPrint(x Parameter) {
	opcode := OpCode{Op: OpPrint, Modes: [3]int{0, 0, x.Mode}}
	c.Program = append(c.Program, Literal(FormatOpCode(opcode)))
	c.Program = append(c.Program, x.Word())
}

// EmitJumpIfTrue writes a new jump-if-true operation with a given set of parameters.
func (c *Compiler) EmitJumpIfTrue(testValue, jumpTo Parameter) {
	opcode := OpCode{Op: OpJumpIfTrue, Modes: [3]int{0, jumpTo.Mode, testValue.Mode}}
	c.Program = append(c.Program, Literal(FormatOpCode(opcode)))
	c.Program = append(c.Program, testValue.Word())
	c.Program = append(c.Program, jumpTo.Word())
}

// EmitJumpIfFalse writes a new jump-if-false operation with a given set of parameters.
func (c *Compiler) EmitJumpIfFalse(testValue, jumpTo Parameter) {
	opcode := OpCode{Op: OpJumpIfFalse, Modes: [3]int{0, jumpTo.Mode, testValue.Mode}}
	c.Program = append(c.Program, Literal(FormatOpCode(opcode)))
	c.Program = append(c.Program, testValue.Word())
	c.Program = append(c.Program, jumpTo.Word())
}

// EmitLessThan writes a new less-than operation with a given set of parameters.
func (c *Compiler) EmitLessThan(a, b, x Parameter) {
	opcode := OpCode{Op: OpLessThan, Modes: [3]int{0, b.Mode, a.Mode}}
	c.Program = append(c.Program, Literal(FormatOpCode(opcode)))
	c.Program = append(c.Program, a.Word())
	c.Program = append(c.Program, b.Word())
	c.Program = append(c.Program, x.Word())
}

// EmitEquals writes a new less-than operation with a given set of parameters.
func (c *Compiler) EmitEquals(a, b, x Parameter) {
	opcode := OpCode{Op: OpEquals, Modes: [3]int{0, b.Mode, a.Mode}}
	c.Program = append(c.Program, Literal(FormatOpCode(opcode)))
	c.Program = append(c.Program, a.Word())
	c.Program = append(c.Program, b.Word())
	c.Program = append(c.Program, x.Word())
}

// EmitSet is a helper that emits an `Add(0, A, X)`, effectively
// setting the value X to A,
func (c *Compiler) EmitSet(a, x Parameter) {
	opcode := OpCode{Op: OpAdd, Modes: [3]int{0, a.Mode, ParameterModeValue}}
	c.Program = append(c.Program, Literal(FormatOpCode(opcode)))
	c.Program = append(c.Program, Literal(0))
	c.Program = append(c.Program, a.Word())
	c.Program = append(c.Program, x.Word())
}

// EmitLiteral emits a literal to the program, returning the index of that literal.
// This typically used to dump values to the program directly that will
// be jumped around.
func (c *Compiler) EmitLiteral(value int) int {
	index := len(c.Program)
	c.Program = append(c.Program, Literal(value))
	return index
}
