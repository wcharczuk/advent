package intcode

import (
	"fmt"
	"testing"

	"github.com/blend/go-sdk/assert"
)

func Test_ParseOpCode(t *testing.T) {
	assert := assert.New(t)

	testCases := [...]struct {
		Input    int
		Expected OpCode
	}{
		{Input: OpAdd, Expected: OpCode{Op: OpAdd, Modes: [3]int{0, 0, 0}}},
		{Input: OpMul, Expected: OpCode{Op: OpMul, Modes: [3]int{0, 0, 0}}},
		{Input: OpHalt, Expected: OpCode{Op: OpHalt, Modes: [3]int{0, 0, 0}}},
		{Input: OpPrint, Expected: OpCode{Op: OpPrint, Modes: [3]int{0, 0, 0}}},
		{Input: 102, Expected: OpCode{Op: OpMul, Modes: [3]int{0, 0, 1}}},
		{Input: 1002, Expected: OpCode{Op: OpMul, Modes: [3]int{0, 1, 0}}},
		{Input: 10002, Expected: OpCode{Op: OpMul, Modes: [3]int{1, 0, 0}}},
		{Input: 11002, Expected: OpCode{Op: OpMul, Modes: [3]int{1, 1, 0}}},
		{Input: 11102, Expected: OpCode{Op: OpMul, Modes: [3]int{1, 1, 1}}},
		{Input: 1008, Expected: OpCode{Op: OpEquals, Modes: [3]int{0, 1, 0}}},
		{Input: 104, Expected: OpCode{Op: OpPrint, Modes: [3]int{0, 0, 1}}},
	}

	for _, tc := range testCases {
		parsed, err := ParseOpCode(tc.Input)
		assert.Nil(err)
		assert.Equal(tc.Expected.Op, parsed.Op, fmt.Sprint(tc.Input))
		assert.Equal(tc.Expected.Modes, parsed.Modes, fmt.Sprint(tc.Input))
	}
}

func Test_ParseOpCode_Regression(t *testing.T) {
	assert := assert.New(t)

	opCode, err := ParseOpCode(104)
	assert.Nil(err)
	assert.Equal(OpPrint, opCode.Op)
	assert.Equal([3]int{0, 0, 1}, opCode.Modes)
	assert.Equal(ParameterModeValue, opCode.Mode(0))
}

func Test_FormatOpCode(t *testing.T) {
	assert := assert.New(t)

	testCases := [...]struct {
		Input    OpCode
		Expected int
	}{
		{Input: OpCode{Op: OpAdd}, Expected: OpAdd},
		{Input: OpCode{Op: OpMul}, Expected: OpMul},
		{Input: OpCode{Op: OpHalt}, Expected: OpHalt},
		{Input: OpCode{Op: OpMul, Modes: [3]int{1, 0, 0}}, Expected: 10002},
		{Input: OpCode{Op: OpMul, Modes: [3]int{0, 1, 0}}, Expected: 1002},
		{Input: OpCode{Op: OpMul, Modes: [3]int{0, 0, 1}}, Expected: 102},
		{Input: OpCode{Op: OpMul, Modes: [3]int{0, 1, 1}}, Expected: 1102},
		{Input: OpCode{Op: OpMul, Modes: [3]int{1, 1, 1}}, Expected: 11102},
	}

	for _, tc := range testCases {
		assert.Equal(tc.Expected, FormatOpCode(tc.Input), fmt.Sprintf("Modes: %#v", tc.Input.Modes))
	}
}

func Test_OpCode_Mode(t *testing.T) {
	assert := assert.New(t)

	oc := OpCode{Modes: [3]int{0, 0, 0}}
	assert.Equal(0, oc.Mode(0))
	assert.Equal(0, oc.Mode(1))
	assert.Equal(0, oc.Mode(2))

	oc = OpCode{Modes: [3]int{1, 0, 0}}
	assert.Equal(0, oc.Mode(0))
	assert.Equal(0, oc.Mode(1))
	assert.Equal(1, oc.Mode(2))

	oc = OpCode{Modes: [3]int{1, 1, 0}}
	assert.Equal(0, oc.Mode(0))
	assert.Equal(1, oc.Mode(1))
	assert.Equal(1, oc.Mode(2))

	oc = OpCode{Modes: [3]int{1, 1, 1}}
	assert.Equal(1, oc.Mode(0))
	assert.Equal(1, oc.Mode(1))
	assert.Equal(1, oc.Mode(2))

	oc = OpCode{Modes: [3]int{0, 1, 1}}
	assert.Equal(1, oc.Mode(0))
	assert.Equal(1, oc.Mode(1))
	assert.Equal(0, oc.Mode(2))

	oc = OpCode{Modes: [3]int{0, 0, 1}}
	assert.Equal(1, oc.Mode(0))
	assert.Equal(0, oc.Mode(1))
	assert.Equal(0, oc.Mode(2))
}

func Test_OpCode_EndToEnd(t *testing.T) {
	assert := assert.New(t)

	ignoreError := func(oc OpCode, err error) OpCode { return oc }

	rawOpCode := 2
	assert.Equal(rawOpCode, FormatOpCode(ignoreError(ParseOpCode(rawOpCode))))

	rawOpCode = 1108
	assert.Equal(rawOpCode, FormatOpCode(ignoreError(ParseOpCode(rawOpCode))))
}
