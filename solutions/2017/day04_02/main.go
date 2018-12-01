package main

import (
	"log"

	"github.com/wcharczuk/advent/pkg/collections"
	"github.com/wcharczuk/advent/pkg/combinatorics"
	"github.com/wcharczuk/advent/pkg/fileutil"
	"github.com/wcharczuk/advent/pkg/stringutil"
)

func main() {
	var valid int
	err := fileutil.ReadByLines("./testdata/input", func(line string) error {
		words := stringutil.SplitOnWhitespace(line)
		lookup := collections.NewSetOfString()
		for _, word := range words {
			for _, anagram := range combinatorics.Anagrams(word) {
				if lookup.Contains(anagram) {
					return nil
				}
			}
			lookup.Add(word)
		}

		valid++
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	println(valid)
}
