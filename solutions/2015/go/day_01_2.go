package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	codeFile := "../testdata/day1"

	index := 0
	floor := 0

	if f, err := os.Open(codeFile); err == nil {
		defer f.Close()

		chunk := make([]byte, 32)
		for {
			readBytes, err := f.Read(chunk)
			if err == io.EOF {
				break
			}
			readData := chunk[:readBytes]
			for _, b := range readData {
				current := string(b)
				index = index + 1
				println(current, index, floor)
				if current == "(" {
					floor = floor + 1
				} else if current == ")" {
					floor = floor - 1
				}
				if floor == -1 {
					break
				}
			}
			if floor == -1 {
				break
			}
		}
	} else {
		fmt.Errorf("%v", err)
	}

	fmt.Printf("Santa arrives at basement at index: %d\n", index)
}
