package intcode

import "fmt"

// InputConstant returns an input handler that returns a constant.
func InputConstant(value int64) func() int64 {
	return func() int64 {
		return value
	}
}

// OutputCapture returns an output capture handler that redirects a value
// to a channel.
func OutputCapture(values chan int64) func(int64) {
	return func(v int64) {
		values <- v
	}
}

// OutputCaptureValue returns a output capture handler that sets a value reference.
func OutputCaptureValue(value *int64) func(int64) {
	return func(v int64) { *value = v }
}

// OutputHandlers is a helper for not having to write the function definition.
func OutputHandlers(handlers ...func(int64)) []func(int64) {
	return handlers
}

// OpWidths returns the total width of a given set of ops.
func OpWidths(ops ...int64) (total int64) {
	for _, op := range ops {
		total += OpWidth(op)
	}
	return
}

// OpWidth returns the width, or number of instructions, per operator.
func OpWidth(op int64) int64 {
	switch op {
	case OpHalt:
		return 1
	case OpAdd:
		return 4
	case OpMul:
		return 4
	case OpInput:
		return 2
	case OpPrint:
		return 2
	case OpJumpIfTrue:
		return 3
	case OpJumpIfFalse:
		return 3
	case OpLessThan:
		return 4
	case OpEquals:
		return 4
	case OpRelativeBase:
		return 2
	default:
		panic(fmt.Sprintf("invalid op for OpWidth: %d", op))
	}
}
