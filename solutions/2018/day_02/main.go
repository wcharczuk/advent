package main

import (
	"github.com/wcharczuk/advent/pkg/fileutil"
	"github.com/wcharczuk/advent/pkg/log"
)

func main() {
	var hasTwo, hasThree int
	err := fileutil.ReadByLines("./input", func(line string) error {
		var hadTwo, hadThree bool
		letterCounts := map[rune]int{}
		for _, letter := range line {
			if count, ok := letterCounts[letter]; ok {
				letterCounts[letter] = count + 1
			} else {
				letterCounts[letter] = 1
			}
		}

		for _, count := range letterCounts {
			if count == 3 && !hadThree {
				hasThree++
				hadThree = true
			}
			if count == 2 && !hadTwo {
				hasTwo++
				hadTwo = true
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Solution(hasTwo * hasThree)
}
