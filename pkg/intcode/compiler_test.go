package intcode

import (
	"testing"

	"github.com/blend/go-sdk/assert"
)

func Test_Compiler(t *testing.T) {
	assert := assert.New(t)

	compiler := new(Compiler)
	compiler.CreateSymbol("limit")
	compiler.CreateSymbol("userInput")

	// store 8 at `limit`
	compiler.SetSymbolValue("limit", 8)

	// read user input into `userInput`
	compiler.Input(
		compiler.SymbolStoreParameter("userInput"),
	)

	compiler.Equals(
		compiler.SymbolReferenceParameter("userInput"), // a == the value at `userInput`
		compiler.SymbolReferenceParameter("limit"),     // b == the value at `limit`
		compiler.PC(0), // store a 1 or 0 at the current PC address
	)

	// print the result of the equals
	compiler.Print(
		compiler.PC(-2), // read from old PC address
	)

	program := compiler.Compile()
	assert.Equal([]int{3, 10, 8, 10, 9, 5, 4, 5, 99, 8, 0}, program) // hard coded check

	var value int
	computer := New(program, OptName("compiler-test"), OptDebug(true))
	computer.InputHandler = InputConstant(8)
	computer.OutputHandlers = OutputHandlers(OutputCaptureValue(&value))
	assert.Nil(computer.Run())
	assert.Equal(1, value)

	computer.InputHandler = InputConstant(7)
	computer.Reset()
	assert.Nil(computer.Run())
	assert.Equal(0, value)
}
