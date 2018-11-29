package stringutil

// CaselessTrimPrefix trims a prefix from a corpus ignoring case.
func CaselessTrimPrefix(corpus, prefix string) string {
	corpusLen := len(corpus)
	prefixLen := len(prefix)

	if corpusLen < prefixLen {
		return corpus
	}

	for x := 0; x < prefixLen; x++ {
		charCorpus := uint(corpus[x])
		charPrefix := uint(prefix[x])

		if charCorpus-LowerA <= lowerDiff {
			charCorpus = charCorpus - 0x20
		}

		if charPrefix-LowerA <= lowerDiff {
			charPrefix = charPrefix - 0x20
		}
		if charCorpus != charPrefix {
			return corpus
		}
	}

	return corpus[prefixLen:]
}
