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
		min, max := util.Math.MinAndMaxOfInt(cols...)
		total += (max - min)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	println("total", total)
}
