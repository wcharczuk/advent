package intcode

import (
	"fmt"
	"io"
)

// New returns a computer for a given program.
func New(program []int64, options ...ComputerOption) *Computer {
	c := &Computer{
		Program: program,
	}
	c.LoadProgram()
	for _, opt := range options {
		opt(c)
	}
	return c
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

	PC      int64
	RB      int64
	Op      OpCode
	A, B, X int64
	Program []int64
	Memory  map[int64]int64

	IsRunning      bool
	InputHandler   func() int64
	OutputHandlers []func(int64)
}

// LoadProgram loads the program into memory.
func (c *Computer) LoadProgram() {
	c.Memory = make(map[int64]int64)
	for index := range c.Program {
		c.Memory[int64(index)] = c.Program[index]
	}
}

// Reset resets all the computer registers to their default values.
func (c *Computer) Reset() {
	c.LoadProgram()
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
	case OpRelativeBase:
		return c.RelativeBase()
	default:
		return fmt.Errorf("%q: %d", ErrInvalidOpCode, c.Op.Op)
	}
}

// Add implements the add operator.
func (c *Computer) Add() error {
	var err error
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
	c.PC = c.PC + OpWidth(OpAdd)
	return nil
}

// Mul implements the multiply operator.
func (c *Computer) Mul() error {
	var err error
	if c.A, _, err = c.Load(1); err != nil {
		return err
	}
	if c.B, _, err = c.Load(2); err != nil {
		return err
	}
	if c.X, _, err = c.LoadForStore(3); err != nil {
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
	c.PC = c.PC + OpWidth(OpInput)
	return nil
}

// Print implements the print operator.
func (c *Computer) Print() error {
	var err error
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
	c.PC = c.PC + OpWidth(OpPrint)
	return nil
}

// JumpIfTrue implements the jump if true operator.
func (c *Computer) JumpIfTrue() error {
	var err error
	if c.A, _, err = c.Load(1); err != nil {
		return err
	}
	if c.B, _, err = c.Load(2); err != nil {
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
	if c.A, _, err = c.Load(1); err != nil {
		return err
	}
	if c.B, _, err = c.Load(2); err != nil {
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
	if c.A, _, err = c.Load(1); err != nil {
		return err
	}
	if c.B, _, err = c.Load(2); err != nil {
		return err
	}
	if c.X, _, err = c.LoadForStore(3); err != nil {
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
	if c.A, _, err = c.Load(1); err != nil {
		return err
	}
	if c.B, _, err = c.Load(2); err != nil {
		return err
	}
	if c.X, _, err = c.LoadForStore(3); err != nil {
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

// RelativeBase sets the current RB value.
func (c *Computer) RelativeBase() error {
	var err error
	if c.A, _, err = c.Load(1); err != nil {
		return err
	}

	c.RB = c.RB + c.A
	c.PC = c.PC + OpWidth(OpRelativeBase)
	return nil
}

// LoadForStore loads a value for a register to be used as a storage address (typically the X  register).
func (c *Computer) LoadForStore(offset int) (result int64, mode int, err error) {
	addr := c.PC + int64(offset)
	var addr2 int64
	if c.Debug {
		defer func() {
			var logEntry string
			if mode == 0 {
				logEntry = fmt.Sprintf("%q (%d+%d) loadmode &%d > %d", c.Name, c.PC, offset, addr, result)
			} else if mode == 2 {
				logEntry = fmt.Sprintf("%q (%d+%d) loadmode %d+(%d) > %d", c.Name, c.PC, offset, c.RB, addr2, result)
			}
			fmt.Fprintln(c.DebugLog, logEntry)
		}()
	}
	if addr < 0 {
		err = fmt.Errorf("%v; address %d", ErrInvalidAddress, addr)
		return
	}

	// get the parameter mode
	mode = c.Op.Mode(offset - 1)
	switch mode {
	case ParameterModeReference: // 0, the default
		result = c.Memory[addr]
		return
	case ParameterModeRelative:
		addr2 = c.Memory[addr]
		result = c.RB + addr2
		return

	default:
		err = fmt.Errorf("invalid parameter mode for store: %d", mode)
		return
	}
}

// Load loads a value from a given offset from the PC.
func (c *Computer) Load(offset int) (result int64, mode int, err error) {
	// the address we're loading from is the PC + an offset.
	addr := c.PC + int64(offset)

	// these are helpers
	var addr2, relativeValue int64

	// debug logs
	if c.Debug {
		defer func() {
			var logEntry string
			if mode == 0 {
				logEntry = fmt.Sprintf("%q (%d+%d) loadmode &%d->&%d > %d", c.Name, c.PC, offset, addr, addr2, result)
			} else if mode == 1 {
				logEntry = fmt.Sprintf("%q (%d+%d) loadmode &%d > %d", c.Name, c.PC, offset, addr, result)
			} else if mode == 2 {
				logEntry = fmt.Sprintf("%q (%d+%d) loadmode %d+(%d) > %d", c.Name, c.PC, offset, c.RB, relativeValue, result)
			}
			fmt.Fprintln(c.DebugLog, logEntry)
		}()
	}

	// check if the load address is valid
	if addr < 0 {
		err = fmt.Errorf("%v; address %d", ErrInvalidAddress, addr)
		return
	}

	// get the parameter mode
	mode = c.Op.Mode(offset - 1)

	switch mode {
	case ParameterModeReference: // 0, the default

		addr2 = c.Memory[addr]
		if addr2 < 0 {
			err = fmt.Errorf("%v; address %d", ErrInvalidAddress, addr2)
			return
		}
		result = c.Memory[addr2]
		return

	case ParameterModeValue:
		result = c.Memory[addr]
		return

	case ParameterModeRelative:
		relativeValue = c.Memory[addr]
		addr2 = c.RB + relativeValue
		if addr2 < 0 {
			err = fmt.Errorf("%v; address %d", ErrInvalidAddress, addr2)
			return
		}
		result = c.Memory[addr2]
		return

	default:
		err = ErrInvalidParameterMode
		return
	}
}

// Store writes a value to the address stored in the X register.
func (c *Computer) Store(value int64) error {
	if c.X < 0 {
		return ErrInvalidAddress
	}
	if c.Debug {
		logEntry := fmt.Sprintf("%q (%d) store &%d < %d", c.Name, c.PC, c.X, value)
		fmt.Fprintln(c.DebugLog, logEntry)
	}
	c.Memory[c.X] = value
	return nil
}
