package array

// TwoOfBool returns a 2d array of ints
func TwoOfBool(height, width int) [][]bool {
	output := make([][]bool, height)
	for y := 0; y < height; y++ {
		output[y] = make([]bool, width)
	}
	return output
}

// TwoOfInt returns a 2d array of ints
func TwoOfInt(height, width int) [][]int {
	output := make([][]int, height)
	for y := 0; y < height; y++ {
		output[y] = make([]int, width)
	}
	return output
}

// TwoOfString returns a 2d array of strings.
func TwoOfString(height, width int) [][]string {
	output := make([][]string, height)
	for y := 0; y < height; y++ {
		output[y] = make([]string, width)
	}
	return output
}
