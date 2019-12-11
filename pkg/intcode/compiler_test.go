package intcode

/*
func Test_Compiler(t *testing.T) {
	assert := assert.New(t)

	compiler := new(Compiler)
	compiler.CreateSymbol("limit")
	compiler.CreateSymbol("userInput")
	compiler.CreateSymbol("temp")

	// store 8 at `limit`
	compiler.HelperEmitSet(
		compiler.Value(8),
		compiler.Value(compiler.Symbol("limit")),
	)

	// read user input into `userInput`
	compiler.EmitInput(
		compiler.Value(compiler.Symbol("userInput")),
	)

	// if userInput == limit, set temp = 1
	compiler.EmitEquals(
		compiler.Reference(compiler.Symbol("userInput")), // a == the value at `userInput`
		compiler.Reference(compiler.Symbol("limit")),     // b == the value at `limit`
		compiler.Value(compiler.Symbol("temp")),
	)

	// temp == 1, jump to halt
	compiler.EmitJumpIfFalse(
		compiler.ReferenceSymbol("temp"),
		compiler.ValuePC(int64(OpWidths(OpPrint, OpJumpIfFalse)+1)), // jump over the print and the jump to halt ...
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
		compiler.ValuePC(int64(OpWidths(OpPrint, OpJumpIfFalse)+1)), // jump over the print and the jump to halt ...
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

	var value int64
	computer := New(program,
		OptName("compiler-test"),
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
*/
