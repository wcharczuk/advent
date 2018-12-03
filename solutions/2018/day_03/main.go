package main

import (
	"github.com/wcharczuk/advent/pkg/fileutil"
	"github.com/wcharczuk/advent/pkg/log"
)

func main() {
	err := fileutil.ReadByLines("./input", func(line string) error {
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Solution("none")
}
