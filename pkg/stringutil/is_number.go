package stringutil

import "strconv"

// IsInteger returns if a string is an integer.
func IsInteger(input string) bool {
	_, err := strconv.Atoi(input)
	return err == nil
}

// IsFloat returns if a string represents a number
func IsFloat(input string) bool {
	_, err := strconv.ParseFloat(input, 64)
	return err == nil
}
