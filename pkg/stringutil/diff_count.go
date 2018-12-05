package stringutil

// DiffCount returns the number of letters that are different between two strings.
// If the strings are different lengths, DiffCount will return -1.
func DiffCount(a, b string) (runesDifferent int) {
	if len(a) != len(b) {
		return -1
	}

	aRunes := []rune(a)
	bRunes := []rune(b)
	strlen := len(aRunes)

	for index := 0; index < strlen; index++ {
		if aRunes[index] != bRunes[index] {
			runesDifferent++
		}
	}
	return
}
