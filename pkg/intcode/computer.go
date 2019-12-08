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
		Program: program,
		Memory:  memory,
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

// OptDebugLog sets the log output collector.
func OptDebugLog(log io.Writer) ComputerOption {
	return func(c *Computer) {
		c.DebugLog = log
	}
}

// Computer is a state machine that processes a program.
type Computer struct {
	Name     string
	Debug    bool
	DebugLog io.Writer

	PC      int
	Op      OpCode
	A, B, X int
	Program []int
	Memory  []int

	IsRunning      bool
	InputHandler   func() int
	OutputHandlers []func(int)
}

// Reset resets all the computer registers to their default values.
func (c *Computer) Reset() {
	c.Memory = c.Program
	c.PC, c.A, c.B, c.X = 0, 0, 0, 0
	c.Op = OpCode{}
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

// Tick applies a computer tick, reading the current op code
// and potentially associated parameters.
func (c *Computer) Tick() error {
	var err error
	c.Op, err = ParseOpCode(c.Memory[c.PC])
	if err != nil {
		return err
	}
	if c.Debug {
		logEntry := fmt.Sprintf("%q (%d) %s", c.Name, c.PC, c.Op.String())
		fmt.Fprintln(c.DebugLog, logEntry)
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
		return fmt.Errorf("%q: %d", ErrInvalidOpCode, c.Op.Op)
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
	c.PC = c.PC + OpWidth(OpAdd)
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
	c.PC = c.PC + OpWidth(OpMul)
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
		value = ReadInt()
	}
	if err = c.Store(value); err != nil {
		return err
	}
	c.PC = c.PC + OpWidth(OpInput)
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
	c.PC = c.PC + OpWidth(OpPrint)
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
	c.PC = c.PC + OpWidth(OpJumpIfTrue)
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
	c.PC = c.PC + OpWidth(OpJumpIfFalse)
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
	c.PC = c.PC + OpWidth(OpLessThan)
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
	c.PC = c.PC + OpWidth(OpEquals)
	return nil
}

// Load loads a value directly without processing the parameter modes.
func (c *Computer) Load(offset int) (result int, err error) {
	addr := c.PC + offset
	if c.Debug {
		defer func() {
			logEntry := fmt.Sprintf("%q (%d+%d) load &%d > %d", c.Name, c.PC, offset, addr, result)
			fmt.Fprintln(c.DebugLog, logEntry)
		}()
	}
	if len(c.Memory) <= addr {
		err = fmt.Errorf("%v; address %d", ErrInvalidAddress, addr)
		return
	}
	result = c.Memory[addr]
	return
}

// LoadMode loads a value from a given offset from the PC.
func (c *Computer) LoadMode(offset int) (result int, mode int, err error) {
	addr := c.PC + offset
	var addr2 int
	if c.Debug {
		defer func() {
			var logEntry string
			if mode == 0 {
				logEntry = fmt.Sprintf("%q (%d+%d) loadmode &%d->&%d > %d", c.Name, c.PC, offset, addr, addr2, result)
			} else {
				logEntry = fmt.Sprintf("%q (%d+%d) loadmode &%d > %d", c.Name, c.PC, offset, addr, result)
			}
			fmt.Fprintln(c.DebugLog, logEntry)
		}()
	}
	if len(c.Memory) <= addr {
		err = fmt.Errorf("%v; address %d", ErrInvalidAddress, addr)
		return
	}
	mode = c.Op.Mode(offset - 1)
	switch mode {
	case ParameterModeReference:
		addr2 = c.Memory[addr]
		if len(c.Memory) <= addr2 {
			err = fmt.Errorf("%v; address %d", ErrInvalidAddress, addr2)
			return
		}
		result = c.Memory[addr2]
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
		logEntry := fmt.Sprintf("%q (%d) store &%d < %d", c.Name, c.PC, c.X, value)
		fmt.Fprintln(c.DebugLog, logEntry)
	}
	c.Memory[c.X] = value
	return nil
}
