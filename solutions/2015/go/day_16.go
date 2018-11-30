package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/blend/go-sdk/util"
)

const ANALYSIS_RESULTS = `children: 3
cats: 7
samoyeds: 2
pomeranians: 3
akitas: 0
vizslas: 0
goldfish: 5
trees: 3
cars: 2
perfumes: 1
`

type Aunt struct {
	Id         int
	Attributes map[string]int
}

func parseEntry(input string) Aunt {
	parts := strings.Split(input, " ")
	a := Aunt{}
	a.Attributes = map[string]int{}
	a.Id, _ = strconv.Atoi(strings.Replace(parts[1], ":", "", 1))

	key := strings.Replace(parts[2], ":", "", 1)
	valueStr := strings.Replace(parts[3], ",", "", 1)
	a.Attributes[key], _ = strconv.Atoi(valueStr)

	key = strings.Replace(parts[4], ":", "", 1)
	valueStr = strings.Replace(parts[5], ",", "", 1)
	a.Attributes[key], _ = strconv.Atoi(valueStr)

	key = strings.Replace(parts[6], ":", "", 1)
	valueStr = strings.Replace(parts[7], ",", "", 1)
	a.Attributes[key], _ = strconv.Atoi(valueStr)

	return a
}

func parseAnalysisResults(results string) map[string]int {
	output := map[string]int{}
	lines := strings.Split(results, "\n")
	for _, line := range lines {
		parts := strings.Split(line, " ")
		if len(parts) > 1 {
			key := strings.Replace(parts[0], ":", "", 1)
			value, _ := strconv.Atoi(parts[1])
			output[key] = value
		}
	}
	return output
}

func matches(aunt Aunt, analysisResults map[string]int) bool {
	for key, value := range aunt.Attributes {
		if expected, hasExpected := analysisResults[key]; hasExpected {
			if expected != value {
				return false
			}
		}
	}
	return true
}

func main() {
	codeFile := "../testdata/day16"

	aunts := []Aunt{}
	util.File.ReadByLines(codeFile, func(line string) error {
		aunts = append(aunts, parseEntry(line))
		return nil
	})

	analyisResults := parseAnalysisResults(ANALYSIS_RESULTS)

	possibleAunts := []Aunt{}
	for _, a := range aunts {
		if matches(a, analyisResults) {
			possibleAunts = append(possibleAunts, a)
		}
	}

	fmt.Println("possible aunts:")
	for _, a := range possibleAunts {
		fmt.Printf("%#v\n", a)
	}
}
