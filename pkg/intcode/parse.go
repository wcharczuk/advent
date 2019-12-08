package intcode

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// ParseString parses a program given as a string.
func ParseString(program string) ([]int, error) {
	return Parse(strings.NewReader(program))
}

// Parse parses a program from a given reader.
func Parse(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	var line string
	var pieces []string
	var err error
	var op, lineNumber, width int
	var a, b, x Parameter
	c := new(Compiler)
	for scanner.Scan() {
		line = scanner.Text()
		if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") {
			// skip empty lines and comments
			continue
		}

		// handle symbols (i.e. .NAME)
		if strings.HasPrefix(line, ".") {
			line = strings.TrimPrefix(line, ".")
			c.CreateSymbol(line)
			continue
		}

		pieces = strings.Split(line, " ")
		op, err = LookupOp(pieces[0])
		if err != nil {
			return nil, fmt.Errorf("invalid program; %q @ line %d", err, lineNumber)
		}
		width = OpWidth(op)
		if len(pieces) < width {
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

		switch op {
		case OpHalt:
			c.EmitHalt()
		case OpAdd:
			c.EmitAdd(a, b, x)
		case OpMul:
			c.EmitMul(a, b, x)
		case OpInput:
			c.EmitInput(a)
		case OpPrint:
			c.EmitPrint(a)
		case OpJumpIfTrue:
			c.EmitJumpIfTrue(a, b)
		case OpJumpIfFalse:
			c.EmitJumpIfFalse(a, b)
		case OpLessThan:
			c.EmitLessThan(a, b, x)
		case OpEquals:
			c.EmitEquals(a, b, x)
		}

	}
	return c.Compile(), nil
}

// ParseParameter parses a single parameter.
// A parameter should be in the form:
//    <value>
//    &<addr>
//    "name"
//    &"name"
//    pc(<offset>)
//    &pc(<offset>)
//    sym(<offset>)
//    &sym(<offset>)
func ParseParameter(c *Compiler, raw string) (param Parameter, err error) {
	var symbol, ok, pc bool
	mode := ParameterModeValue
	if strings.HasPrefix(raw, "&") {
		mode = ParameterModeReference
		raw = strings.TrimPrefix(raw, "&")
	}
	if strings.HasPrefix(raw, "pc(") {
		raw = strings.TrimPrefix(raw, "pc(")
		if !strings.HasSuffix(raw, ")") {
			err = fmt.Errorf("parse parameter; pc param is missing closing ')'")
			return
		}
		raw = strings.TrimSuffix(raw, ")")
		pc = true
	} else if strings.HasPrefix(raw, "sym(") {
		raw = strings.TrimPrefix(raw, "sym(")
		if !strings.HasSuffix(raw, ")") {
			err = fmt.Errorf("parse parameter; sym param is missing closing ')'")
			return
		}
		raw = strings.TrimSuffix(raw, ")")
		symbol = true
	}

	var value int
	if strings.HasPrefix(raw, ".") {
		raw = strings.TrimPrefix(raw, ".")
		symbol = true
		value, ok = c.SymbolAddrs[raw]
		if !ok {
			err = fmt.Errorf("parse parameter; unknown symbol %q", raw)
		}
	} else {
		value, err = strconv.Atoi(raw)
		if err != nil {
			return
		}
	}

	param.Mode = mode
	param.Symbol = symbol
	param.PC = pc
	param.Value = value
	return
}

// LookupOp translates an op name to an op code.
func LookupOp(opName string) (int, error) {
	switch strings.ToLower(opName) {
	case "halt":
		return OpHalt, nil
	case "add":
		return OpAdd, nil
	case "mul":
		return OpMul, nil
	case "input":
		return OpInput, nil
	case "print":
		return OpPrint, nil
	case "jump-if-true":
		return OpJumpIfTrue, nil
	case "jump-if-false":
		return OpJumpIfFalse, nil
	case "less-than":
		return OpLessThan, nil
	case "equals":
		return OpEquals, nil
	default:
		return -1, fmt.Errorf("lookup op; invalid op name: %s", opName)
	}
}
