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
	var first, current, last, total int
	for i, c := range contents {
		current, err = strconv.Atoi(string(c))
		if err != nil {
			log.Fatal(err)
		}

		if i == 0 {
			first = current
		}

		if i > 0 && last == current {
			total += current
		}

		last = current
	}

	if first == last {
		total += first
	}

	println("total", total)
}
