package intcode

// Parameter modes
const (
	ParameterModeReference = 0
	ParameterModeValue     = 1
	ParameterModeRelative  = 2
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
