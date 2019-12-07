package intcode

import (
	"fmt"
	"io"
)

// New returns a computer for a given program.
func New(program []int, options ...ComputerOption) *Computer {
	memory := make([]int, len(program))
	copy(memory, program)
	c := Computer{
		Memory: memory,
	}
	for _, opt := range options {
		opt(&c)
	}
	return &c
}

// ComputerOption is a mutator for computers.
type ComputerOption func(*Computer)

// OptName sets the computer name.
func OptName(name string) ComputerOption {
	return func(c *Computer) {
		c.Name = name
	}
}

// OptDebug sets if we should show debug output.
func OptDebug(debug bool) ComputerOption {
	return func(c *Computer) {
		c.Debug = debug
	}
}

// Computer is a state machine that processes a program.
type Computer struct {
	Name  string
	Debug bool
	Log   []string

	PC      int
	Op      OpCode
	A, B, X int
	Memory  []int

	IsRunning      bool
	InputHandler   func() int
	OutputHandlers []func(int)
}

// Run runs the program.
func (c *Computer) Run() (err error) {
	defer func() {
		c.IsRunning = false
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()

	c.IsRunning = true
	for {
		err = c.Tick()
		if err == ErrHalt {
			return nil
		}
		if err != nil {
			return err
		}
	}
}

// WriteLogTo writes log contents to a given writer.
func (c *Computer) WriteLogTo(w io.Writer) (err error) {
	for _, entry := range c.Log {
		_, err = io.WriteString(w, entry+"\n")
		if err != nil {
			return
		}
	}
	return
}

// Tick applies a computer tick, reading the current op code
// and potentially associated parameters.
func (c *Computer) Tick() error {
	var err error
	c.Op, err = ParseOpCode(c.Memory[c.PC])
	if err != nil {
		return err
	}
	if c.Debug {
		c.Log = append(c.Log, fmt.Sprintf("%q (%d) %s", c.Name, c.PC, c.Op.String()))
	}

	switch c.Op.Op {
	case OpHalt:
		return ErrHalt
	case OpAdd:
		return c.Add()
	case OpMul:
		return c.Mul()
	case OpInput: // input
		return c.Input()
	case OpPrint: // print
		return c.Print()
	case OpJumpIfTrue: // jump if true
		return c.JumpIfTrue()
	case OpJumpIfFalse: // jump if false
		return c.JumpIfFalse()
	case OpLessThan: // less than
		return c.LessThan()
	case OpEquals: // equals
		return c.Equals()
	default:
		return ErrInvalidOpCode
	}
}

// Add implements the add operator.
func (c *Computer) Add() error {
	var err error
	if c.A, _, err = c.LoadMode(1); err != nil {
		return err
	}
	if c.B, _, err = c.LoadMode(2); err != nil {
		return err
	}
	if c.X, err = c.Load(3); err != nil {
		return err
	}
	if err = c.Store(c.A + c.B); err != nil {
		return err
	}
	c.PC = c.PC + 4
	return nil
}

// Mul implements the multiply operator.
func (c *Computer) Mul() error {
	var err error
	if c.A, _, err = c.LoadMode(1); err != nil {
		return err
	}
	if c.B, _, err = c.LoadMode(2); err != nil {
		return err
	}
	if c.X, err = c.Load(3); err != nil {
		return err
	}
	if err = c.Store(c.A * c.B); err != nil {
		return err
	}
	c.PC = c.PC + 4
	return nil
}

// Input implements the input operator.
func (c *Computer) Input() error {
	var err error
	if c.X, err = c.Load(1); err != nil {
		return err
	}
	var value int
	if c.InputHandler != nil {
		value = c.InputHandler()
	} else {
		value = readInt()
	}
	if err = c.Store(value); err != nil {
		return err
	}
	c.PC = c.PC + 2
	return nil
}

// Print implements the print operator.
func (c *Computer) Print() error {
	var err error
	if c.X, _, err = c.LoadMode(1); err != nil {
		return err
	}
	if len(c.OutputHandlers) > 0 {
		for _, handler := range c.OutputHandlers {
			handler(c.X)
		}
	} else {
		fmt.Println(c.X)
	}
	c.PC = c.PC + 2
	return nil
}

// JumpIfTrue implements the jump if true operator.
func (c *Computer) JumpIfTrue() error {
	var err error
	if c.A, _, err = c.LoadMode(1); err != nil {
		return err
	}
	if c.B, _, err = c.LoadMode(2); err != nil {
		return err
	}
	if c.A > 0 {
		c.PC = c.B
		return nil
	}
	c.PC = c.PC + 3
	return nil
}

// JumpIfFalse implements the jump if false operator.
func (c *Computer) JumpIfFalse() error {
	var err error
	if c.A, _, err = c.LoadMode(1); err != nil {
		return err
	}
	if c.B, _, err = c.LoadMode(2); err != nil {
		return err
	}
	if c.A == 0 {
		c.PC = c.B
		return nil
	}
	c.PC = c.PC + 3
	return nil
}

// LessThan implements the less than operator.
func (c *Computer) LessThan() error {
	var err error
	if c.A, _, err = c.LoadMode(1); err != nil {
		return err
	}
	if c.B, _, err = c.LoadMode(2); err != nil {
		return err
	}
	if c.X, err = c.Load(3); err != nil {
		return err
	}
	if c.A < c.B {
		c.Store(1)
	} else {
		c.Store(0)
	}
	c.PC = c.PC + 4
	return nil
}

// Equals implements the equals operator.
func (c *Computer) Equals() error {
	var err error
	if c.A, _, err = c.LoadMode(1); err != nil {
		return err
	}
	if c.B, _, err = c.LoadMode(2); err != nil {
		return err
	}
	if c.X, err = c.Load(3); err != nil {
		return err
	}
	if c.A == c.B {
		c.Store(1)
	} else {
		c.Store(0)
	}
	c.PC = c.PC + 4
	return nil
}

// Load loads a value directly without processing the parameter modes.
func (c *Computer) Load(offset int) (result int, err error) {
	addr := c.PC + offset
	if c.Debug {
		defer func() {
			c.Log = append(c.Log, fmt.Sprintf("%q (%d) load %d &%d > %d", c.Name, c.PC, offset, addr, result))
		}()
	}
	if len(c.Memory) <= addr {
		err = ErrInvalidAddress
		return
	}
	result = c.Memory[addr]
	return
}

// LoadMode loads a value from a given offset from the PC.
func (c *Computer) LoadMode(offset int) (result int, mode int, err error) {
	addr := c.PC + offset
	if c.Debug {
		defer func() {
			c.Log = append(c.Log, fmt.Sprintf("%q (%d) loadmode %d &%d (%v) > %d", c.Name, c.PC, offset, addr, ParameterMode(mode), result))
		}()
	}
	if len(c.Memory) <= addr {
		err = ErrInvalidAddress
		return
	}
	mode = c.Op.Mode(offset - 1)
	switch mode {
	case ParameterModeReference:
		addr = c.Memory[addr]
		if len(c.Memory) <= addr {
			err = ErrInvalidAddress
			return
		}
		result = c.Memory[addr]
		return
	case ParameterModeValue:
		result = c.Memory[addr]
		return
	default:
		err = ErrInvalidParameterMode
		return
	}
}

// Store writes a value to the address stored in the X register.
func (c *Computer) Store(value int) error {
	if len(c.Memory) <= c.X {
		return ErrInvalidAddress
	}
	if c.Debug {
		c.Log = append(c.Log, fmt.Sprintf("%q (%d) store &%d < %d", c.Name, c.PC, c.X, value))
	}
	c.Memory[c.X] = value
	return nil
}

func readInt() int {
	fmt.Print("Input: ")
	var i int
	fmt.Scanf("%d", &i)
	return i
}
