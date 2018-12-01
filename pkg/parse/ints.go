package parse

import (
	"strconv"
)

// Ints parses a list of strings as ints.
func Ints(values ...string) (output []int, err error) {
	output = make([]int, len(values))
	var parsed int64
	for index, value := range values {
		parsed, err = strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil, err
		}
		output[index] = int(parsed)
	}
	return
}
