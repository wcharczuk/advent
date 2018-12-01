package mathutil

// MinAndMaxOfInt returns both the min and max in one pass.
func MinAndMaxOfInt(values ...int) (min, max int) {
	if len(values) == 0 {
		return
	}
	min = values[0]
	max = values[0]
	for _, v := range values {
		if max < v {
			max = v
		}
		if min > v {
			min = v
		}
	}
	return
}

// MinAndMax returns both the min and max in one pass.
func MinAndMax(values ...float64) (min, max float64) {
	if len(values) == 0 {
		return
	}
	min = values[0]
	max = values[0]
	for _, v := range values {
		if max < v {
			max = v
		}
		if min > v {
			min = v
		}
	}
	return
}
