package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/blend/go-sdk/collections"
	"github.com/blend/go-sdk/util"
)

const (
	START    = "e"
	END      = "CRnCaCaCaSiRnBPTiMgArSiRnSiRnMgArSiRnCaFArTiTiBSiThFYCaFArCaCaSiThCaPBSiThSiThCaCaPTiRnPBSiThRnFArArCaCaSiThCaSiThSiRnMgArCaPTiBPRnFArSiThCaSiRnFArBCaSiRnCaPRnFArPMgYCaFArCaPTiTiTiBPBSiThCaPTiBPBSiRnFArBPBSiRnCaFArBPRnSiRnFArRnSiRnBFArCaFArCaCaCaSiThSiThCaCaPBPTiTiRnFArCaPTiBSiAlArPBCaCaCaCaCaSiRnMgArCaSiThFArThCaSiThCaSiRnCaFYCaSiRnFYFArFArCaSiRnFYFArCaSiRnBPMgArSiThPRnFArCaSiRnFArTiRnSiRnFYFArCaSiRnBFArCaSiRnTiMgArSiThCaSiThCaFArPRnFArSiRnFArTiTiTiTiBCaCaSiRnCaCaFYFArSiThCaPTiBPTiBCaSiThSiRnMgArCaF"
	TEST_END = "HOHOHO"
)

type xform [2]string

type byReplacement []xform

func (br byReplacement) Len() int           { return len(br) }
func (br byReplacement) Swap(i, j int)      { br[i], br[j] = br[j], br[i] }
func (br byReplacement) Less(i, j int) bool { return br[i][1] > br[j][1] }

func main() {
	xfs := []xform{}
	util.File.ReadByLines("../testdata/day19", func(line string) error {
		parts := strings.Split(line, " => ")
		xfs = append(xfs, xform{parts[0], parts[1]})
		return nil
	})

	results := find(xfs, START, END)
	fmt.Printf("shortest path is of length %d\n", len(results))
}

func find(xforms []xform, start, end string) []xform {
	return findImpl(xforms, start, end, []xform{})
}

type state struct {
	Start string
	Path  []xform
}

func findImpl(xforms []xform, start, end string, path []xform) []xform {
	if end == start {
		return path
	}

	sorted_xforms := sortReplacementSizeDescending(xforms)

	for _, xf := range sorted_xforms {
		results := reverseApply(xf, end)
		newPath := dupepath(path)
		newPath = append(newPath, xf)

		if results != nil {
			for _, result := range results {
				if result == start {
					return append(path, xf)
				}
				recursive := findImpl(xforms, start, result, newPath)
				if recursive != nil {
					return recursive
				}
			}
		}
	}

	return nil
}

func sortReplacementSizeDescending(input []xform) []xform {
	newBuffer := make([]xform, len(input))
	copy(newBuffer, input)
	sort.Sort(byReplacement(newBuffer))
	return newBuffer
}

func contains(xf xform, path []xform) bool {
	for _, pathItem := range path {
		if xf[0] == pathItem[0] && xf[1] == pathItem[1] {
			return true
		}
	}
	return false
}

func reverseApply(xf xform, input string) []string {
	ss := collections.SetOfString{}
	key, replacement := xf[0], xf[1]
	inputBytes := []byte(input)
	for x := 0; x < len(inputBytes); x++ {
		var cursor string
		if x+len(replacement) < len(inputBytes) {
			cursor = string(slice(inputBytes, x, x+len(replacement)))
		} else {
			cursor = string(slice(inputBytes, x, len(inputBytes)))
		}

		if cursor == replacement {
			pre := slice(inputBytes, 0, x)
			post := slice(inputBytes, x+len(replacement), len(inputBytes))
			joined := append(pre, []byte(key)...)
			final := append(joined, post...)
			ss.Add(string(final))
		}
	}
	return ss.ToArray()
}

func apply(xf xform, input string) []string {
	ss := collections.StringSet{}
	key, replacement := xf[0], xf[1]
	inputBytes := []byte(input)
	for x := 0; x < len(inputBytes); x++ {
		var cursor string
		if x+len(key) < len(inputBytes) {
			cursor = string(slice(inputBytes, x, x+len(key)))
		} else {
			cursor = string(slice(inputBytes, x, len(inputBytes)))
		}

		if cursor == key {
			pre := slice(inputBytes, 0, x)
			post := slice(inputBytes, x+len(key), len(inputBytes))
			joined := append(pre, []byte(replacement)...)
			final := append(joined, post...)
			ss.Add(string(final))
		}
	}
	return ss.ToArray()
}

func slice(buffer []byte, start, end int) []byte {
	output := []byte{}
	for x := start; x < end; x++ {
		output = append(output, buffer[x])
	}
	return output
}

func dupepath(path []xform) []xform {
	newCopy := make([]xform, len(path))
	copy(newCopy, path)
	return newCopy
}
