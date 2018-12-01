package main

import (
	"strconv"

	"github.com/wcharczuk/advent/pkg/fileutil"
	"github.com/wcharczuk/advent/pkg/log"
)

func main() {
	var freq int64
	err := fileutil.ReadByLines("./input", func(line string) error {
		adjustment, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			return err
		}
		freq = freq + adjustment
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Context("solution").Print(freq)
}
