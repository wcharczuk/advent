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
