package parse

import "strconv"

// Floats parses a list of strings as ints.
func Floats(values ...string) (output []float64, err error) {
	output = make([]float64, len(values))
	var parsed float64
	for index, value := range values {
		parsed, err = strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, err
		}
		output[index] = parsed
	}
	return
}
