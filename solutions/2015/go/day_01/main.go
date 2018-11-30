package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	codeFile := "../../testdata/day1"

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
				if string(b) == "(" {
					floor = floor + 1
				} else if string(b) == ")" {
					floor = floor - 1
				}
			}
		}
	} else {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	fmt.Printf("Santa arrives at floor: %d\n", floor)
}
