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
var testProgram = []int{
	3, 21, // input => &21 [0, 1]
	1008, 21, 8, 20, // equals &21 8 => &20 [2, 3, 4, 5]
	1005, 20, 22, // jump-if-true &20 to 22 [6, 7, 8]
	107, 8, 21, 20, // less-than 8 &21 => &20 [9, 10, 11, 12]
	1006, 20, 31, // jump-if-false &20 to 31 [13, 14, 15]
	1106, 0, 36, // jump-if-false 0 to 36 [16, 17, 18] // i.e. definitely jump to 36
	98, 0, 0, // <variables/unknown> [19, 20, 21]
	1002, 21, 125, 20, // mul &21-> 125 => &20 [22, 23, 24, 25]
	4, 20, // print &20 [26, 27]
	1105, 1, 46, // jump-if-true 1 to 46 [28, 29, 30]
	104, 999, // print 999 [31, 32]
	1105, 1, 46, // jump-if-true 1 to 46 [33, 34, 35]
	1101, 1000, 1, 20, // add 1000 1 => &20 [36, 37, 38, 39]
	4, 20, // print &20 [40, 41]
	1105, 1, 46, // jump-if-true 1 46 [42, 43, 44]
	98, // <variable> [45] // why is this here ??
	99, // HALT [46]
}

// testProgramAltered is the program from day 5 that returns 999 if
// the number is less than 8, 1000 if the number equals 8, and 1001
// if the number is greater than 8, but without an erroneous .45
var testProgramAltered = []int{
	3, 21, // input => &21 [0, 1]
	1008, 21, 8, 20, // equals &21 8 => &20 [2, 3, 4, 5]
	1005, 20, 22, // jump-if-true &20 to 22 [6, 7, 8]
	107, 8, 21, 20, // less-than 8 &21 => &20 [9, 10, 11, 12]
	1006, 20, 31, // jump-if-false &20 to 31 [13, 14, 15]
	1106, 0, 36, // jump-if-false 0 to 36 [16, 17, 18] // i.e. definitely jump to 36
	98, 0, 0, // <variables/unknown> [19, 20, 21]
	1002, 21, 125, 20, // mul &21-> 125 => &20 [22, 23, 24, 25]
	4, 20, // print &20 [26, 27]
	1105, 1, 45, // jump-if-true 1 to 46 [28, 29, 30]
	104, 999, // print 999 [31, 32]
	1105, 1, 45, // jump-if-true 1 to 46 [33, 34, 35]
	1101, 1000, 1, 20, // add 1000 1 => &20 [36, 37, 38, 39]
	4, 20, // print &20 [40, 41]
	1105, 1, 45, // jump-if-true 1 46 [42, 43, 44]
	99, // HALT [45]
}

func Test_Computer_Run(t *testing.T) {
	assert := assert.New(t)

	runTest := func(input int) (int, error) {
		computer := New(testProgramAltered, OptName("run_test"))
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
