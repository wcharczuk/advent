package profanity

// Error is a string error.
type Error string

// Error implements `error`
func (e Error) Error() string { return string(e) }

// Errors
const (
	ErrFailure Error = "profanity failure"
)
