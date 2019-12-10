package intcode

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
