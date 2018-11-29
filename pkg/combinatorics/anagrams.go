package combinatorics

// Anagrams is a form of permutations that is of a fixed length (i.e. order matters).
// It is very similar to permutations of string but uses word inputs instead of individual strings.
func Anagrams(word string) []string {
	if len(word) <= 1 {
		return []string{word}
	}

	output := []string{}
	var letter byte
	var pre []byte
	var post []byte
	var joined []byte
	for x := 0; x < len(word); x++ {
		workingWord := make([]byte, len(word))
		copy(workingWord, []byte(word))
		letter = workingWord[x]
		pre = workingWord[0:x]
		post = workingWord[x+1 : len(word)]
		joined = append(pre, post...)
		for _, subResult := range Anagrams(string(joined)) {
			output = append(output, string(letter)+subResult)
		}
	}
	return output
}
