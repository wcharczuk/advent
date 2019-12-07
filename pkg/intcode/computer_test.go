package intcode

import (
	"testing"

	"github.com/blend/go-sdk/assert"
)

func Test_Computer_LoadMode(t *testing.T) {
	assert := assert.New(t)

	program := []int{1008, 2, 10}
	computer := New(program)
	result, mode, err := computer.LoadMode(1)
	assert.Nil(err)
	assert.Equal(ParameterModeReference, mode)
	assert.Equal(10, result)
}

// testProgram is the program from day 5 that returns 999 if
// the number is less than 8, 1000 if the number equals 8, and 1001
// if the number is greater than 8.
var testProgram = []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}

func Test_Computer_Run(t *testing.T) {
	assert := assert.New(t)

	runTest := func(input int) (int, error) {
		computer := New(testProgram, OptName("run_test"))
		computer.InputHandler = func() int {
			return input
		}
		var output int
		computer.OutputHandlers = append(computer.OutputHandlers, func(value int) {
			output = value
		})
		err := computer.Run()
		return output, err
	}

	result, err := runTest(5)
	assert.Nil(err)
	assert.Equal(999, result)

	result, err = runTest(8)
	assert.Nil(err)
	assert.Equal(1000, result)

	result, err = runTest(9)
	assert.Nil(err)
	assert.Equal(1001, result)
}
