package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	AND = iota
	OR  = iota

	LSHIFT = iota
	RSHIFT = iota

	NOT = iota
)

type wire interface {
	Signal() uint16
}

type input struct {
	Id    string
	Value uint16
}

func (i input) Signal() uint16 {
	return i.Value
}

type direct struct {
	Id           string
	Input        wire
	InputId      string
	cachedResult *uint16
}

func (d *direct) Signal() uint16 {
	if d.cachedResult != nil {
		return *d.cachedResult
	}
	result := d.Input.Signal()
	d.cachedResult = &result
	return result
}

type unaryGate struct {
	Id           string
	Input        wire
	InputId      string
	Op           byte
	cachedResult *uint16
}

func (u *unaryGate) Signal() uint16 {
	if u.cachedResult != nil {
		return *u.cachedResult
	}
	result := ^u.Input.Signal()
	u.cachedResult = &result
	return result
}

type shift struct {
	Id           string
	Input        wire
	InputId      string
	Offset       uint
	Op           byte
	cachedResult *uint16
}

func (s *shift) Signal() uint16 {
	if s.cachedResult != nil {
		return *s.cachedResult
	}

	result := uint16(0)
	if s.Op == LSHIFT {
		result = s.Input.Signal() << s.Offset
	} else if s.Op == RSHIFT {
		result = s.Input.Signal() >> s.Offset
	}
	s.cachedResult = &result
	return result
}

type binaryGate struct {
	Id       string
	A, B     wire
	AId, BId string
	Op       byte

	cachedResult *uint16
}

func (bg *binaryGate) Signal() uint16 {
	if bg.cachedResult != nil {
		return *bg.cachedResult
	}

	result := uint16(0)
	if bg.Op == AND {
		result = bg.A.Signal() & bg.B.Signal()
	} else if bg.Op == OR {
		result = bg.A.Signal() | bg.B.Signal()
	}
	bg.cachedResult = &result
	return result
}

func parseInstruction(instruction string, wires map[string]wire) {
	instructionParts := strings.Split(instruction, " ")

	if value, parseErr := strconv.Atoi(instructionParts[0]); parseErr == nil && len(instructionParts) < 4 {
		w := input{}
		w.Value = uint16(value)
		w.Id = instructionParts[2]
		wires[w.Id] = &w
		return
	}

	if instructionParts[0] == "NOT" {
		w := unaryGate{}
		w.Op = NOT
		w.Id = instructionParts[3]
		w.InputId = instructionParts[1]

		if inputWire, hasInputWire := wires[instructionParts[1]]; hasInputWire {
			w.Input = inputWire
			w.Id = instructionParts[3]
		}

		wires[w.Id] = &w
		return
	}

	opName := instructionParts[1]
	if opName == "LSHIFT" || opName == "RSHIFT" {
		w := shift{}
		w.Id = instructionParts[4]
		w.InputId = instructionParts[0]

		if opName == "LSHIFT" {
			w.Op = LSHIFT
		} else if opName == "RSHIFT" {
			w.Op = RSHIFT
		}
		value, _ := strconv.Atoi(instructionParts[2])
		w.Offset = uint(value)

		if inputWire, hasInputWire := wires[instructionParts[0]]; hasInputWire {
			w.Input = inputWire
		}

		wires[w.Id] = &w
		return
	}

	if len(instructionParts) < 5 {
		w := direct{}
		w.Id = instructionParts[2]
		w.InputId = instructionParts[0]

		if input, hasInput := wires[instructionParts[0]]; hasInput {
			w.Input = input
		}
		wires[w.Id] = &w
		return
	}

	w := binaryGate{}
	w.Id = instructionParts[4]

	if instructionParts[1] == "AND" {
		w.Op = AND
	} else if instructionParts[1] == "OR" {
		w.Op = OR
	}

	a := instructionParts[0]
	b := instructionParts[2]

	if signalValue, signalErr := strconv.Atoi(a); signalErr == nil {
		w.A = &input{Value: uint16(signalValue)}
		w.AId = "<input>"
	} else {
		w.AId = a
		if a, hasA := wires[a]; hasA {
			w.A = a
		}
	}
	if signalValue, signalErr := strconv.Atoi(b); signalErr == nil {
		w.B = &input{Value: uint16(signalValue)}
		w.BId = "<input>"
	} else {
		w.BId = b
		if b, hasB := wires[b]; hasB {
			w.B = b
		}
	}

	wires[w.Id] = &w
	return
}

func fixWireReferences(wires map[string]wire) {
	for _, w := range wires {
		if typed, isTyped := w.(*direct); isTyped {
			if typed.Input == nil {
				typed.Input = wires[typed.InputId]
			}
		} else if typed, isTyped := w.(*unaryGate); isTyped {
			if typed.Input == nil {
				typed.Input = wires[typed.InputId]
			}
		} else if typed, isTyped := w.(*shift); isTyped {
			if typed.Input == nil {
				typed.Input = wires[typed.InputId]
			}
		} else if typed, isTyped := w.(*binaryGate); isTyped {
			if typed.A == nil {
				typed.A = wires[typed.AId]
			}
			if typed.B == nil {
				typed.B = wires[typed.BId]
			}
		}
	}
}

func checkWireReferences(wires map[string]wire) bool {
	ok := true
	for _, w := range wires {
		ok = ok && checkWireReference(w, wires)
	}
	return ok
}

func checkWireReference(w wire, wires map[string]wire) bool {
	if typed, isTyped := w.(*direct); isTyped {
		if typed.Input == nil {
			println(typed.Id, "has bad input:", typed.InputId)
			if wireExists(typed.InputId, wires) {
				println(typed.InputId, "exists in wires reference.")

			} else {
				println(typed.InputId, "doesnt exist in wires reference.")
			}
			return false
		}
	} else if typed, isTyped := w.(*unaryGate); isTyped {
		if typed.Input == nil {
			println(typed.Id, "has bad input:", typed.InputId)
			if wireExists(typed.InputId, wires) {
				println(typed.InputId, "exists in wires reference.")

			} else {
				println(typed.InputId, "doesnt exist in wires reference.")
			}
			return false
		}
	} else if typed, isTyped := w.(*shift); isTyped {
		if typed.Input == nil {
			println(typed.Id, "has bad input:", typed.InputId)
			if wireExists(typed.InputId, wires) {
				println(typed.InputId, "exists in wires reference.")

			} else {
				println(typed.InputId, "doesnt exist in wires reference.")
			}
			return false
		}
	} else if typed, isTyped := w.(*binaryGate); isTyped {
		if typed.A == nil {
			println(typed.Id, "has bad input:", typed.AId)
			if wireExists(typed.AId, wires) {
				println(typed.AId, "exists in wires reference.")

			} else {
				println(typed.AId, "doesnt exist in wires reference.")
			}
			return false
		}
		if typed.B == nil {
			println(typed.Id, "has bad input:", typed.BId)
			if wireExists(typed.BId, wires) {
				println(typed.BId, "exists in wires reference.")

			} else {
				println(typed.BId, "doesnt exist in wires reference.")
			}
			return false
		}
	} else if _, isTyped := w.(*input); !isTyped {
		fmt.Printf("%v could't be type coerced\n", w)
		return false
	}
	return true
}

func wireExists(wireId string, wires map[string]wire) bool {
	_, exists := wires[wireId]
	return exists
}

func main() {
	dataFile := "../testdata/day7"

	wires := map[string]wire{}
	if f, err := os.Open(dataFile); err == nil {
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			instruction := scanner.Text()
			parseInstruction(instruction, wires)
		}
	}

	checkCycles := 0
	maxCycles := len(wires)
	for !checkWireReferences(wires) && checkCycles <= maxCycles {
		fixWireReferences(wires)
		checkCycles++
	}

	targetWire := wires["a"]
	fmt.Println("Target Wire a:", targetWire.Signal())
}
