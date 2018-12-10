package log

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// a standard singleton.
var std Std

// Std writes to stdout or stderr.
type Std struct {
	Contexts []string
}

// Context returns a new context.
func (l Std) Context(label string) Logger {
	return Std{Contexts: append(l.Contexts, label)}
}

// Contextf returns a new context by format.
func (l Std) Contextf(format string, args ...interface{}) Logger {
	return Std{Contexts: append(l.Contexts, fmt.Sprintf(format, args...))}
}

// Timestamp returns the current time as a string.
func (l Std) Timestamp() string {
	return ColorLightBlack.Apply(time.Now().UTC().Format(time.RFC3339))
}

// Heading returns the current contexts as a string.
func (l Std) Heading() string {
	return strings.Join(l.Contexts, " > ")
}

// Fprint writes to a given writer.
func (l Std) Fprint(w io.Writer, args ...interface{}) {
	fmt.Fprint(w, l.Timestamp())
	fmt.Fprint(w, " ")
	if heading := l.Heading(); heading != "" {
		fmt.Fprintf(w, " [%s] ", heading)
	}
	fmt.Fprint(w, args...)
	fmt.Fprintln(w)
}

// Print prints to stdout.
func (l Std) Print(args ...interface{}) {
	l.Fprint(os.Stdout, args...)
}

// Printf prints to stdout.
func (l Std) Printf(format string, args ...interface{}) {
	l.Fprint(os.Stdout, fmt.Sprintf(format, args...))
}

// Error prints to stderr.
func (l Std) Error(args ...interface{}) {
	l.Fprint(os.Stderr, args...)
}

// Errorf prints to stderr.
func (l Std) Errorf(format string, args ...interface{}) {
	l.Fprint(os.Stderr, fmt.Sprintf(format, args...))
}
