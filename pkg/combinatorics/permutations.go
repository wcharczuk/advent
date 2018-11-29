package combinatorics

// PermutationsOfInt returns the possible orderings of the values array.
// Use "permutations" when order matters.
func PermutationsOfInt(values ...int) [][]int {
	if len(values) == 1 {
		return [][]int{values}
	}

	output := [][]int{}
	for x := 0; x < len(values); x++ {
		workingValues := make([]int, len(values))
		copy(workingValues, values)
		value := workingValues[x]
		pre := workingValues[0:x]
		post := workingValues[x+1 : len(values)]

		joined := append(pre, post...)

		for _, inner := range PermutationsOfInt(joined...) {
			output = append(output, append([]int{value}, inner...))
		}
	}

	return output
}

// PermutationsOfFloat returns the possible orderings of the values array.
// Use "permutations" when order matters.
func PermutationsOfFloat(values ...float64) [][]float64 {
	if len(values) == 1 {
		return [][]float64{values}
	}

	output := [][]float64{}
	for x := 0; x < len(values); x++ {
		workingValues := make([]float64, len(values))
		copy(workingValues, values)
		value := workingValues[x]
		pre := workingValues[0:x]
		post := workingValues[x+1 : len(values)]

		joined := append(pre, post...)

		for _, inner := range PermutationsOfFloat(joined...) {
			output = append(output, append([]float64{value}, inner...))
		}
	}

	return output
}

// PermutationsOfString returns the possible orderings of the values array (i.e. when order matters).
// Note: Use "combinations" when order doesn't matter.
func PermutationsOfString(values ...string) [][]string {
	if len(values) == 1 {
		return [][]string{values}
	}

	output := [][]string{}
	for x := 0; x < len(values); x++ {
		workingValues := make([]string, len(values))
		copy(workingValues, values)
		value := workingValues[x]
		pre := workingValues[0:x]
		post := workingValues[x+1 : len(values)]
		joined := append(pre, post...)
		for _, inner := range PermutationsOfString(joined...) {
			output = append(output, append([]string{value}, inner...))
		}
	}

	return output
}
