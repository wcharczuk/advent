package intcode

import (
	"fmt"
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
	Name    string
	Debug   bool
	LogItem OpLog
	Log     []OpLog

	PC      int
	Current OpCode
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

// Tick applies a computer tick, reading the current op code
// and potentially associated parameters.
func (c *Computer) Tick() error {
	c.Current = ParseOpCode(c.Memory[c.PC])

	c.LogItem = OpLog{Name: c.Name, Op: c.Current, PC: c.PC}
	if c.Debug {
		fmt.Println(c.LogItem.String())
	}
	defer func() {
		if c.Debug {
			fmt.Println(c.LogItem.String())
		}
		c.Log = append(c.Log, c.LogItem)
	}()

	switch c.Current.Op {
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
func (c *Computer) Load(offset int) (int, error) {
	addr := c.PC + offset
	if len(c.Memory) <= addr {
		return 0, ErrInvalidAddress
	}
	return c.Memory[addr], nil
}

// LoadMode loads a value from a given offset from the PC.
func (c *Computer) LoadMode(offset int) (result int, mode int, err error) {
	addr := c.PC + offset
	if len(c.Memory) <= addr {
		return 0, 0, ErrInvalidAddress
	}
	mode = c.Current.Mode(offset - 1)
	defer func() {
		c.LogItem.Parameters = append(c.LogItem.Parameters, OpLogParameter{
			IsReference: mode == 0,
			Addr:        addr,
			Value:       result,
		})
	}()
	switch mode {
	case 0:
		addr = c.Memory[addr]
		if len(c.Memory) <= addr {
			return 0, 0, ErrInvalidAddress
		}
		result = c.Memory[addr]
		return
	case 1:
		result = c.Memory[addr]
		return
	default:
		err = ErrInvalidParameterMode
		return
	}
}

// Store writes a value to the address stored in the X register.
func (c *Computer) Store(value int) error {
	defer func() {
		c.LogItem.Store = OpLogParameter{
			IsReference: true,
			Addr:        c.X,
			Value:       value,
		}
	}()
	if len(c.Memory) <= c.X {
		return ErrInvalidAddress
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
