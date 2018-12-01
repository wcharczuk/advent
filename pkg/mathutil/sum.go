package mathutil

// Sum adds all the numbers of a slice together
func Sum(input ...float64) (sum float64) {
	if len(input) == 0 {
		return
	}
	for _, n := range input {
		sum += n
	}
	return
}

// SumInts adds all the numbers of a slice together
func SumInts(input ...int) (sum int) {
	if len(input) == 0 {
		return
	}
	for _, n := range input {
		sum += n
	}
	return
}
