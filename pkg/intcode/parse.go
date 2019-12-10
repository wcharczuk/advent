package intcode

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// ParseString parses a program given as a string.
func ParseString(program string) ([]int64, error) {
	return Parse(strings.NewReader(program))
}

// Parse parses a program from a given reader.
func Parse(r io.Reader) ([]int64, error) {
	scanner := bufio.NewScanner(r)
	var line string
	var pieces []string
	var err error
	var ok bool
	var instruction Instruction
	var width int
	var lineNumber int
	var a, b, x Parameter
	c := new(Compiler)
	for scanner.Scan() {
		line = scanner.Text()
		if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") {
			// skip empty lines and comments
			continue
		}

		pieces = strings.Split(line, " ")
		instruction, ok = LookupOp(pieces[0])
		if !ok {
			return nil, fmt.Errorf("invalid program; %q @ line %d", err, lineNumber)
		}
		width = instruction.Width()
		if len(pieces) < int(width) {
			return nil, fmt.Errorf("invalid program; instruction does not have required number of arguments (%d), %q @ line %d", width, line, lineNumber)
		}
		if width > 3 {
			a, err = ParseParameter(c, pieces[1])
			if err != nil {
				return nil, fmt.Errorf("invalid program; %q @ line %d", err, lineNumber)
			}
			b, err = ParseParameter(c, pieces[2])
			if err != nil {
				return nil, fmt.Errorf("invalid program; %q @ line %d", err, lineNumber)
			}
			x, err = ParseParameter(c, pieces[3])
			if err != nil {
				return nil, fmt.Errorf("invalid program; %q @ line %d", err, lineNumber)
			}
		} else if width > 2 {
			a, err = ParseParameter(c, pieces[1])
			if err != nil {
				return nil, fmt.Errorf("invalid program; %q @ line %d", err, lineNumber)
			}
			b, err = ParseParameter(c, pieces[2])
			if err != nil {
				return nil, fmt.Errorf("invalid program; %q @ line %d", err, lineNumber)
			}
		} else if width > 1 {
			a, err = ParseParameter(c, pieces[1])
			if err != nil {
				return nil, fmt.Errorf("invalid program; %q @ line %d", err, lineNumber)
			}
		}

	}
	return c.Compile(), nil
}

// ParseParameter parses a single parameter.
// A parameter should be in the form:
//    <value>
//    &<addr>
//    @<offset>
func ParseParameter(c *Compiler, raw string) (param Parameter, err error) {
	mode := ParameterModeValue
	if strings.HasPrefix(raw, "&") {
		mode = ParameterModeReference
		raw = strings.TrimPrefix(raw, "&")
	} else if strings.HasPrefix(raw, "@") {
		mode = ParameterModeRelative
		raw = strings.TrimPrefix(raw, "@")
	}

	var value int64
	value, err = strconv.ParseInt(raw, 10, 64)
	if err != nil {
		return
	}
	param.Mode = mode
	param.Value = value
	return
}

// LookupOp translates an op name to an op code.
func LookupOp(opName string) (Instruction, bool) {
	instruction, ok := InstructionNames[opName]
	return instruction, ok
}
