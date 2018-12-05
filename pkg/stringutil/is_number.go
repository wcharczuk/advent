package stringutil

import "strconv"

// IsInt returns if a string is an integer.
func IsInt(input string) bool {
	_, err := strconv.Atoi(input)
	return err == nil
}

// IsFloat64 returns if a string represents a number
func IsFloat64(input string) bool {
	_, err := strconv.ParseFloat(input, 64)
	return err == nil
}
