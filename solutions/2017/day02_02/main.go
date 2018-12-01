package main

import (
	"github.com/wcharczuk/advent/pkg/combinatorics"
	"github.com/wcharczuk/advent/pkg/fileutil"
	"github.com/wcharczuk/advent/pkg/log"
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

		pairs := combinatorics.PairsOfInt(cols...)

		for _, p := range pairs {
			if p[0]%p[1] == 0 {
				total += (p[0] / p[1])
			} else if p[1]%p[0] == 0 {
				total += (p[1] / p[0])
			}
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	println("total", total)
}
