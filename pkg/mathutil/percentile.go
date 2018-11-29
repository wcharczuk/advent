package mathutil

import "math"

// Percentile finds the relative standing in a slice of floats.
// `percent` should be given on the interval [0,100.0).
func Percentile(input []float64, percent float64) float64 {
	if len(input) == 0 {
		return 0
	}

	return PercentileOfSorted(SortCopy(input), percent)
}

// PercentileOfSorted finds the relative standing in a sorted slice of floats.
// `percent` should be given on the interval [0,100.0).
func PercentileOfSorted(sortedInput []float64, percent float64) float64 {
	index := (percent / 100.0) * float64(len(sortedInput))
	percentile := float64(0)
	i := int(math.RoundToEven(index))
	if index == float64(int64(index)) {
		percentile = (sortedInput[i-1] + sortedInput[i]) / 2.0
	} else {
		percentile = sortedInput[i-1]
	}

	return percentile
}
