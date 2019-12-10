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
