package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	UINT_MAX      = ^uint(0)
	INT_MAX       = int(UINT_MAX >> 1)
	INT_MIN_VALUE = ^0
)

type stringSet map[string]int

func (ss stringSet) Add(entry string) {
	if _, hasEntry := ss[entry]; !hasEntry {
		ss[entry] = len(ss)
	}
}

func (ss stringSet) IndexOf(entry string) int {
	if index, hasEntry := ss[entry]; hasEntry {
		return index
	} else {
		return -1
	}
}

func (ss stringSet) Contains(entry string) bool {
	if _, hasEntry := ss[entry]; hasEntry {
		return true
	} else {
		return false
	}
}

func (ss stringSet) Remove(entry string) bool {
	if _, hasEntry := ss[entry]; hasEntry {
		delete(ss, entry)
		return true
	}
	return false
}

func (ss stringSet) Len() int {
	return len(ss)
}

func (ss stringSet) Copy() stringSet {
	newSet := stringSet{}
	for key, _ := range ss {
		newSet.Add(key)
	}
	return newSet
}

func (ss stringSet) ToArray() []string {
	output := []string{}
	for key, _ := range ss {
		output = append(output, key)
	}
	return output
}

func (ss stringSet) String() string {
	return strings.Join(ss.ToArray(), ", ")
}

func parseEntry(input string) (string, string, int) {
	inputParts := strings.Split(input, " ")

	from := inputParts[0]
	to := inputParts[2]
	distanceStr := inputParts[4]

	distance, _ := strconv.Atoi(distanceStr)
	return from, to, distance
}

func main() {
	dataFile := "../testdata/day9"

	locations := stringSet{}
	mat := [20][20]int{}

	if f, err := os.Open(dataFile); err == nil {
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			entry := scanner.Text()
			from, to, distance := parseEntry(entry)
			locations.Add(from)
			locations.Add(to)
			fromIndex := locations.IndexOf(from)
			toIndex := locations.IndexOf(to)
			mat[fromIndex][toIndex] = distance
			mat[toIndex][fromIndex] = distance
		}
	}

	allPermutations := permute(locations.ToArray())
	maxDistance := -1
	bestPath := []string{}
	for _, path := range allPermutations {
		distance := 0
		for x := 0; x < len(path)-1; x++ {
			from := path[x]
			to := path[x+1]
			distance += mat[locations.IndexOf(from)][locations.IndexOf(to)]
		}

		if distance > maxDistance {
			maxDistance = distance
			bestPath = path
		}
	}

	fmt.Println("Distance:", maxDistance)
	fmt.Println("Path:", strings.Join(bestPath, ", "))
}

func permute(values []string) [][]string {
	if len(values) == 1 {
		return [][]string{values}
	}

	output := [][]string{}
	for x := 0; x < len(values); x++ {
		value := values[x]
		pre := arraySub(values, 0, x)
		post := arraySub(values, x+1, len(values))

		joined := append(pre, post...)

		for _, inner := range permute(joined) {
			output = append(output, append([]string{value}, inner...))
		}
	}

	return output
}

func arraySub(input []string, from, to int) []string {
	output := []string{}
	for x := from; x < to; x++ {
		output = append(output, input[x])
	}
	return output
}
