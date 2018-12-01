package main

import (
	"log"

	"github.com/wcharczuk/advent/pkg/collections"
	"github.com/wcharczuk/advent/pkg/fileutil"
	"github.com/wcharczuk/advent/pkg/stringutil"
)

func main() {
	var valid int
	err := fileutil.ReadByLines("./testdata/input", func(line string) error {
		words := stringutil.SplitOnWhitespace(line)
		lookup := collections.NewSetOfString()
		for _, word := range words {
			if !lookup.Contains(word) {
				lookup.Add(word)
			} else {
				return nil
			}
		}
		valid++
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	println(valid)
}
