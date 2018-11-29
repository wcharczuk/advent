package main

import (
	"io/ioutil"
	"log"
	"strconv"
)

func main() {
	contents, err := ioutil.ReadFile("./testdata/input")
	if err != nil {
		log.Fatal(err)
	}

	// build the values buffer
	var current int
	values := make([]int, len(contents))
	for i, c := range contents {
		current, err = strconv.Atoi(string(c))
		if err != nil {
			log.Fatal(err)
		}
		values[i] = current
	}

	valueCount := len(values)
	valueCount2 := valueCount >> 1
	var total int
	for i := 0; i < valueCount; i++ {
		if values[i] == values[(i+valueCount2)%valueCount] {
			total += values[i]
		}
	}

	println("total", total)
}
