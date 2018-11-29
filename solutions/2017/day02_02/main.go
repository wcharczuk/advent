package main

import (
	"log"

	"github.com/blendlabs/go-util"
)

func main() {
	var total int
	err := util.File.ReadByLines("./testdata/input", func(line string) error {
		cols, err := util.Parse.Ints(util.String.SplitOnSpace(line)...)
		if err != nil {
			return err
		}

		pairs := util.Combinatorics.PairsOfInt(cols...)

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
