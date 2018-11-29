package combinatorics

// CombinationsOfFloat returns the "power set" of values less the empty set.
// Use "combinations" when the order of the resulting sets do not matter.
func CombinationsOfFloat(values ...float64) [][]float64 {
	possibleValues := PowOfInt(2, uint(len(values))) //less the empty entry
	output := make([][]float64, possibleValues-1)

	for x := 0; x < possibleValues-1; x++ {
		row := []float64{}
		for i := 0; i < len(values); i++ {
			y := 1 << uint(i)
			if y&x == 0 && y != x {
				row = append(row, values[i])
			}
		}
		if len(row) > 0 {
			output[x] = row
		}
	}
	return output
}

// CombinationsOfInt returns the "power set" of values less the empty set.
// Use "combinations" when the order of the resulting sets do not matter.
func CombinationsOfInt(values ...int) [][]int {
	possibleValues := PowOfInt(2, uint(len(values))) //less the empty entry
	output := make([][]int, possibleValues-1)

	for x := 0; x < possibleValues-1; x++ {
		row := []int{}
		for i := 0; i < len(values); i++ {
			y := 1 << uint(i)
			if y&x == 0 && y != x {
				row = append(row, values[i])
			}
		}
		if len(row) > 0 {
			output[x] = row
		}
	}
	return output
}

// CombinationsOfString returns the "power set" of values less the empty set.
// Use "combinations" when the order of the resulting sets do not matter.
func CombinationsOfString(values ...string) [][]string {
	possibleValues := PowOfInt(2, uint(len(values))) //less the empty entry
	output := make([][]string, possibleValues-1)

	for x := 0; x < possibleValues-1; x++ {
		row := []string{}
		for i := 0; i < len(values); i++ {
			y := 1 << uint(i)
			if y&x == 0 && y != x {
				row = append(row, values[i])
			}
		}
		if len(row) > 0 {
			output[x] = row
		}
	}
	return output
}
