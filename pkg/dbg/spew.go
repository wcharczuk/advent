package dbg

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

// Spew calls `Fspew(os.Stderr, obj)`.
func Spew(obj interface{}, extra ...interface{}) error {
	return Fspew(os.Stderr, obj, extra...)
}

// Fspew writes an object to a given writer, including the file of the caller, the stack
// and some other metadata.
func Fspew(out io.Writer, obj interface{}, extra ...interface{}) error {
	stackTrace := CallersFromStartDepth(DefaultStackTraceDepth)

	take := 3                   // we want to take this many frames
	if len(stackTrace) < take { // but the stacktrace might not be that long
		take = len(stackTrace) // so take what we can from the trace
	}
	var frameSegments []string
	var frame Frame
	for _, framePtr := range stackTrace[0 : take-1] { // `take-1` == always skip proc.go
		frame = Frame(framePtr)
		frameSegments = append([]string{frame.Func()}, frameSegments...)
	}
	if _, err := fmt.Fprintf(out, "%s ", strings.Join(frameSegments, "|")); err != nil {
		return err
	}
	if err := json.NewEncoder(out).Encode(obj); err != nil {
		return err
	}

	if len(extra) > 0 {
		if _, err := fmt.Fprintln(out, extra...); err != nil {
			return err
		}
	}
	return nil
}
