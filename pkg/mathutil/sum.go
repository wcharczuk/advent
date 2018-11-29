package mathutil

// Sum adds all the numbers of a slice together
func Sum(input ...float64) float64 {
	if len(input) == 0 {
		return 0
	}

	sum := float64(0)

	// Add em up
	for _, n := range input {
		sum += n
	}

	return sum
}
