package intcode

// Error is an error
type Error string

// Error implements error.
func (e Error) Error() string { return string(e) }

// Error constants.
const (
	ErrHalt                 Error = "program halted"
	ErrInvalidParameterMode Error = "invalid parameter mode"
	ErrInvalidAddress       Error = "invalid address"
	ErrInvalidOpCode        Error = "invalid op code"
)

// Op is an opcode.
type Op int

// Op values
const (
	OpAdd         = 1
	OpMul         = 2
	OpInput       = 3
	OpPrint       = 4
	OpJumpIfTrue  = 5
	OpJumpIfFalse = 6
	OpLessThan    = 7
	OpEquals      = 8
	OpHalt        = 99
)
