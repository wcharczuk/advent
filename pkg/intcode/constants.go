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

// ParameterMode is a parameter mode.
type ParameterMode int

// String returns the string representation of the mode.
func (pm ParameterMode) String() string {
	switch pm {
	case ParameterModeReference:
		return "ref"
	case ParameterModeValue:
		return "val"
	default:
		return "unknown"
	}
}

// Parameter modes
const (
	ParameterModeReference = 0
	ParameterModeValue     = 1
	ParameterModeRelative  = 2
)

// Op values
const (
	OpAdd          = 1
	OpMul          = 2
	OpInput        = 3
	OpPrint        = 4
	OpJumpIfTrue   = 5
	OpJumpIfFalse  = 6
	OpLessThan     = 7
	OpEquals       = 8
	OpRelativeBase = 9
	OpHalt         = 99
)
