package stringutil

import (
	"strings"
	"unicode"
)

// CompressWhitespace compresses whitespace characters into single spaces.
// It trims leading and trailing whitespace as well.
func CompressWhitespace(text string) (output string) {
	if len(text) == 0 {
		return
	}

	var state int
	for _, r := range text {
		switch state {
		case 0: // non-whitespace
			if unicode.IsSpace(r) {
				state = 1
			} else {
				output = output + string(r)
			}
		case 1: // whitespace
			if !unicode.IsSpace(r) {
				output = output + " " + string(r)
				state = 0
			}
		}
	}

	output = strings.TrimSpace(output)
	return
}
