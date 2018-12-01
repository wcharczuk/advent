package main

import (
	"log"
	"strconv"

	"github.com/wcharczuk/advent/pkg/fileutil"
)

func main() {
	var ops []int
	err := fileutil.ReadByLines("./testdata/input", func(line string) error {
		val, err := strconv.Atoi(line)
		if err != nil {
			return err
		}
		ops = append(ops, val)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	var jc, pc, steps int
	for pc >= 0 && pc < len(ops) {
		jc = ops[pc]
		ops[pc]++
		pc += jc
		steps++
	}

	println("steps", steps)
}
