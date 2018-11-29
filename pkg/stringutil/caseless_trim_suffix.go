package stringutil

// CaselessTrimSuffix trims a case insensitive suffix from a corpus.
func CaselessTrimSuffix(corpus, suffix string) string {
	corpusLen := len(corpus)
	suffixLen := len(suffix)

	if corpusLen < suffixLen {
		return corpus
	}

	for x := 0; x < suffixLen; x++ {
		charCorpus := uint(corpus[corpusLen-(x+1)])
		charSuffix := uint(suffix[suffixLen-(x+1)])

		if charCorpus-LowerA <= lowerDiff {
			charCorpus = charCorpus - 0x20
		}

		if charSuffix-LowerA <= lowerDiff {
			charSuffix = charSuffix - 0x20
		}

		if charCorpus != charSuffix {
			return corpus
		}
	}
	return corpus[:corpusLen-suffixLen]
}
