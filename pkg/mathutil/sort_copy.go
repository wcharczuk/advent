package mathutil

import "sort"

// Copy copies an array of float64s.
func Copy(input []float64) []float64 {
	output := make([]float64, len(input))
	copy(output, input)
	return output
}

// SortCopy copies and sorts an array of floats.
func SortCopy(input []float64) []float64 {
	inputCopy := make([]float64, len(input))
	copy(inputCopy, input)
	sort.Float64s(inputCopy)
	return inputCopy
}
