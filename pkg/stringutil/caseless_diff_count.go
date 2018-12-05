package stringutil

// CaselessDiffCount compares two strings regardless of case and returns the number of runes different.
func CaselessDiffCount(a, b string) (runesDifferent int) {
	aLen := len(a)
	bLen := len(b)
	if aLen != bLen {
		return -1
	}

	var charA, charB uint
	for x := 0; x < aLen; x++ {
		charA = uint(a[x])
		charB = uint(b[x])

		if charA-LowerA <= lowerDiff {
			charA = charA - 0x20
		}
		if charB-LowerA <= lowerDiff {
			charB = charB - 0x20
		}
		if charA != charB {
			runesDifferent++
		}
	}
	return
}
