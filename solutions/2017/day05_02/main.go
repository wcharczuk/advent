package main

import (
	"log"
	"strconv"

	util "github.com/blendlabs/go-util"
)

func main() {
	var ops []int
	err := util.File.ReadByLines("./testdata/input", func(line string) error {
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
		if jc >= 3 {
			ops[pc]--
		} else {
			ops[pc]++
		}
		pc += jc
		steps++
	}

	println("steps", steps)
}
