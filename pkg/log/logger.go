package log

import "os"

// Logger is a type that prints to logs.
type Logger interface {
	Print(args ...interface{})
	Printf(format string, args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})

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

// Solution prints a message and exit(0)s the process.
func Solution(args ...interface{}) {
	std.Context("solution").Print(args...)
	os.Exit(0)
}

// Solutionf prints a message and exit(0)s the process.
func Solutionf(format string, args ...interface{}) {
	std.Context("solution").Printf(format, args...)
	os.Exit(0)
}

// Fatal prints an error message and exit(1)s the process.
func Fatal(args ...interface{}) {
	std.Context("fatal").Error(args...)
	os.Exit(1)
}

// Fatalf prints an error message and exit(1)s the process.
func Fatalf(format string, args ...interface{}) {
	std.Context("fatal").Errorf(format, args...)
	os.Exit(1)
}
