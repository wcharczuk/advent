package combinatorics

// PairsOfFloat64 returns unordered pairs of integers from an array.
func PairsOfFloat64(values ...float64) [][2]float64 {
	if len(values) == 0 {
		return nil
	}

	var output [][2]float64
	for i := 0; i < len(values); i++ {
		for j := i + 1; j < len(values); j++ {
			output = append(output, [2]float64{values[i], values[j]})
		}
	}

	return output
}

// PairsOfInt returns unordered pairs of integers from an array.
func PairsOfInt(values ...int) [][2]int {
	if len(values) == 0 {
		return nil
	}

	var output [][2]int
	for i := 0; i < len(values); i++ {
		for j := i + 1; j < len(values); j++ {
			output = append(output, [2]int{values[i], values[j]})
		}
	}

	return output
}

// PairsOfString returns unordered pairs of strings from an array.
func PairsOfString(values ...string) [][2]string {
	if len(values) == 0 {
		return nil
	}

	var output [][2]string
	for i := 0; i < len(values); i++ {
		for j := i + 1; j < len(values); j++ {
			output = append(output, [2]string{values[i], values[j]})
		}
	}

	return output
}
