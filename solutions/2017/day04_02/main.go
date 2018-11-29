package main

import (
	"log"

	util "github.com/blendlabs/go-util"
	"github.com/blendlabs/go-util/collections"
)

func main() {
	var valid int
	err := util.File.ReadByLines("./testdata/input", func(line string) error {
		words := util.String.SplitOnSpace(line)
		lookup := collections.NewSetOfString()
		for _, word := range words {
			for _, anagram := range util.Combinatorics.Anagrams(word) {
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
