package parse

import (
	"regexp"
)

// MustExtract runs a capture expression on a corpus and panics if there is an error.
// The first result (index 0) will be the corpus if the expression matches
// capture groups start at index 1.
func MustExtract(expr, corpus string) []string {
	results, err := Extract(expr, corpus)
	if err != nil {
		panic(err)
	}
	return results
}

// Extract extracts matches for a given expression.
// The first result (index 0) will be the corpus if the expression matches
// capture groups start at index 1.
func Extract(expr, corpus string) ([]string, error) {
	compiled, err := regexp.Compile(expr)
	if err != nil {
		return nil, err
	}
	return ExtractCompiled(compiled, corpus), nil
}

// ExtractCompiled extracts values from a compiled expression for a corpus.
// The first result (index 0) will be the corpus if the expression matches
// capture groups start at index 1.
func ExtractCompiled(re *regexp.Regexp, corpus string) (matches []string) {
	resultSets := re.FindAllStringSubmatch(corpus, -1)
	for _, resultSet := range resultSets {
		for _, result := range resultSet {
			matches = append(matches, result)
		}
	}

	return
}
