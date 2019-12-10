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

	if instruction, ok := Instructions[c.Op.Op]; ok {
		if c.Debug {
			logEntry := fmt.Sprintf("%q (%d) %s", c.Name, c.PC, instruction.Name())
			fmt.Fprintln(c.DebugLog, logEntry)
		}
		if err = instruction.Action(c); err != nil {
			return err
		}
		if typed, ok := instruction.(InstructionPC); ok {
			if err = typed.MovePC(c); err != nil {
				return nil
			}
		} else {
			c.PC = c.PC + int64(instruction.Width())
		}
	} else {
		return fmt.Errorf("%q: %d", ErrInvalidOpCode, c.Op.Op)
	}
	return nil
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
