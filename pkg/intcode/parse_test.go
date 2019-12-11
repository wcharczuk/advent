package intcode

/*
func Test_ParseString(t *testing.T) {
	assert := assert.New(t)

	assembly := `
# declare locals
.limit
.userInput
.temp
# start the program
add 0 8 .limit
input .userInput
equals &.userInput &.limit .temp
jump-if-false &.temp pc(6)
print 0
# sym(-1) here means jump to halt
jump-if-false 0 sym(-1)
less-than &.userInput &.limit .temp
jump-if-false &.temp pc(6)
print -1
# sym(-1) here means jump to halt
jump-if-false 0 sym(-1)
print 1`

	program, err := ParseString(assembly)
	assert.Nil(err)

	var value int64
	computer := New(program,
		OptName("parser-test"),
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
