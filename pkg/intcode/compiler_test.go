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
	compiler.CreateSymbol("temp")

	// store 8 at `limit`
	compiler.EmitSet(
		compiler.Value(8),
		compiler.StoreSymbol("limit"),
	)

	// read user input into `userInput`
	compiler.EmitInput(
		compiler.StoreSymbol("userInput"),
	)

	// if userInput == limit, set temp = 1
	compiler.EmitEquals(
		compiler.ReferenceSymbol("userInput"), // a == the value at `userInput`
		compiler.ReferenceSymbol("limit"),     // b == the value at `limit`
		compiler.StoreSymbol("temp"),          // store a 1 or 0 at the current PC address
	)

	// temp == 1, jump to halt
	compiler.EmitJumpIfFalse(
		compiler.ReferenceSymbol("temp"),
		compiler.ValuePC(OpWidths(OpPrint, OpJumpIfFalse)+1), // jump over the print and the jump to halt ...
	)
	// print 0 (equals) //
	compiler.EmitPrint(
		compiler.Value(0),
	)
	// goto halt
	compiler.EmitJumpIfFalse(
		compiler.Value(0),
		compiler.ValueHaltAddr(),
	)

	compiler.EmitLessThan(
		compiler.ReferenceSymbol("userInput"), // a == the value at `userInput`
		compiler.ReferenceSymbol("limit"),     // b == the value at `limit`
		compiler.StoreSymbol("temp"),          // store a 1 or 0 at the current PC address
	)
	compiler.EmitJumpIfFalse(
		compiler.ReferenceSymbol("temp"),
		compiler.ValuePC(OpWidths(OpPrint, OpJumpIfFalse)+1), // jump over the print and the jump to halt ...
	)
	compiler.EmitPrint(
		compiler.Value(-1),
	)
	compiler.EmitJumpIfFalse(
		compiler.Value(0),
		compiler.ValueHaltAddr(),
	)

	compiler.EmitPrint(
		compiler.Value(1),
	)
	program := compiler.Compile()

	var value int
	computer := New(program,
		OptName("compiler-test"),
		//OptDebug(true),
		//OptDebugLog(os.Stdout),
	)
	computer.InputHandler = InputConstant(8)
	computer.OutputHandlers = OutputHandlers(OutputCaptureValue(&value))
	assert.Nil(computer.Run())
	assert.Equal(0, value)

	computer.InputHandler = InputConstant(7)
	computer.Reset()
	assert.Nil(computer.Run())
	assert.Equal(-1, value)

	computer.InputHandler = InputConstant(9)
	computer.Reset()
	assert.Nil(computer.Run())
	assert.Equal(1, value)
}
