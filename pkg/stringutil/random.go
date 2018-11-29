package stringutil

import (
	"math/rand"
	"time"
)

var (
	provider = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// RandomLetters returns a new random string composed of letters from the `letters` collection.
func RandomLetters(length int) string {
	return RandomRunes(Letters, length)
}

// RandomNumbers returns a random string of chars from the `numbers` collection.
func RandomNumbers(length int) string {
	return RandomRunes(Numbers, length)
}

// RandomLettersAndNumbers returns a random string composed of chars from the `lettersAndNumbers` collection.
func RandomLettersAndNumbers(length int) string {
	return RandomRunes(LettersAndNumbers, length)
}

// RandomLettersAndNumbersAndSymbols returns a random string composed of chars from the `lettersNumbersAndSymbols` collection.
func RandomLettersAndNumbersAndSymbols(length int) string {
	return RandomRunes(LettersNumbersAndSymbols, length)
}

// RandomRunes returns a random selection of runes from the set.
func RandomRunes(runeset []rune, length int) string {
	runes := make([]rune, length)
	for index := range runes {
		runes[index] = runeset[provider.Intn(len(runeset))]
	}
	return string(runes)
}

// CombineRunsets combines given runsets into a single runset.
func CombineRunsets(runesets ...[]rune) []rune {
	output := []rune{}
	for _, set := range runesets {
		output = append(output, set...)
	}
	return output
}
