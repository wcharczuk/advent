package mathutil

// Mean gets the average of a slice of numbers
func Mean(input []float64) float64 {
	if len(input) == 0 {
		return 0
	}

	sum := Sum(input...)
	return sum / float64(len(input))
}
