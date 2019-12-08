package intcode

// InputConstant returns an input handler that returns a constant.
func InputConstant(value int) func() int {
	return func() int {
		return value
	}
}

// OutputCapture returns an output capture handler that redirects a value
// to a channel.
func OutputCapture(values chan int) func(int) {
	return func(v int) {
		values <- v
	}
}

// OutputCaptureValue returns a output capture handler that sets a value reference.
func OutputCaptureValue(value *int) func(int) {
	return func(v int) { *value = v }
}

// OutputHandlers is a helper for not having to write the function definition.
func OutputHandlers(handlers ...func(int)) []func(int) {
	return handlers
}
