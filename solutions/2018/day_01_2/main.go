package main

import (
	"os"
	"strconv"

	"github.com/wcharczuk/advent/pkg/collections"
	"github.com/wcharczuk/advent/pkg/fileutil"
	"github.com/wcharczuk/advent/pkg/log"
)

func main() {
	seen := collections.NewSetOfInt()
	var freq int
	for {
		err := fileutil.ReadByLines("./input", func(line string) error {
			adjustment, err := strconv.Atoi(line)
			if err != nil {
				return err
			}
			freq = freq + adjustment
			if seen.Contains(freq) {
				log.Print("duplicate: ", freq)
				os.Exit(0)
			}
			seen.Add(freq)
			return nil
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}
