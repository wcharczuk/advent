package main

import (
	"io/ioutil"

	"github.com/wcharczuk/advent/pkg/log"
)

func main() {
	contents, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	//contents := []byte("dabAcCaCBAcCcaDA")

	var ok bool
	for {
		if contents, ok = process(contents); !ok {
			break
		}
	}

	log.Solutionf("%d", len(contents))
	//log.Solutionf("%s %d", contents, len(contents))
}

func process(contents []byte) (output []byte, ok bool) {
	if len(contents) == 0 {
		return
	}
	for index := 0; index < len(contents); index++ {
		if index < len(contents)-1 {
			if matches(contents[index], contents[index+1]) {
				ok = true
				index++
				continue
			}
		}

		output = append(output, contents[index])
	}
	return
}

const (
	// LowerA is the ascii int value for 'a'
	LowerA uint = uint('a')
	// LowerZ is the ascii int value for 'z'
	LowerZ uint = uint('z')
)

var (
	lowerDiff = (LowerZ - LowerA)
)

func matches(a, b byte) bool {
	if !same(a, b) {
		return false
	}
	return a != b
}

func same(a, b byte) bool {
	charA := uint(a)
	charB := uint(b)
	if charA-LowerA <= lowerDiff {
		charA = charA - 0x20
	}
	if charB-LowerA <= lowerDiff {
		charB = charB - 0x20
	}
	return charA == charB
}
