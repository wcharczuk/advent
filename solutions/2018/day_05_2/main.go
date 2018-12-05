package main

import (
	"io/ioutil"
	"unicode"

	"github.com/wcharczuk/advent/pkg/log"
)

func main() {
	contents, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	elements := uniq(contents)

	min := len(contents)
	for _, element := range elements {
		working := less(element, contents)
		var ok bool
		for {
			if working, ok = process(working); !ok {
				break
			}
		}
		if len(working) < min {
			min = len(working)
		}
	}

	log.Solutionf("%d", min)
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

func uniq(contents []byte) []byte {
	values := map[byte]bool{}
	for _, c := range contents {
		values[byte(unicode.ToLower(rune(c)))] = true
	}
	var output []byte
	for value := range values {
		output = append(output, value)
	}
	return output
}

func less(element byte, contents []byte) []byte {
	var output []byte
	for _, c := range contents {
		if !same(element, c) {
			output = append(output, c)
		}
	}
	return output
}

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
