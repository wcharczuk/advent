package log

// Logger is a type that prints to logs.
type Logger interface {
	Print(args ...interface{})
	Printf(format string, args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})

	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})

	// Should nest to a new context
	Context(label string) Logger
}

// Context returns the logger with a new context.
func Context(label string) Logger {
	return std.Context(label)
}

// Print a message
func Print(args ...interface{}) {
	std.Print(args...)
}

// Printf a message
func Printf(format string, args ...interface{}) {
	std.Printf(format, args...)
}

// Error prints an error message
func Error(args ...interface{}) {
	std.Error(args...)
}

// Errorf prints an error message
func Errorf(format string, args ...interface{}) {
	std.Errorf(format, args...)
}

// Fatal prints an error message
func Fatal(args ...interface{}) {
	std.Fatal(args...)
}

// Fatalf prints an error message
func Fatalf(format string, args ...interface{}) {
	std.Fatalf(format, args...)
}
