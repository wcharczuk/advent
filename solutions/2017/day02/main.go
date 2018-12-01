package main

import (
	"log"

	"github.com/wcharczuk/advent/pkg/fileutil"
	"github.com/wcharczuk/advent/pkg/mathutil"
	"github.com/wcharczuk/advent/pkg/parse"
	"github.com/wcharczuk/advent/pkg/stringutil"
)

func main() {
	var total int
	err := fileutil.ReadByLines("./testdata/input", func(line string) error {
		cols, err := parse.Ints(stringutil.SplitOnWhitespace(line)...)
		if err != nil {
			return err
		}
		min, max := mathutil.MinAndMaxOfInt(cols...)
		total += (max - min)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	println("total", total)
}
