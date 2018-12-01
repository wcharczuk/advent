package main

import (
	"fmt"
	"strings"

	"github.com/wcharczuk/advent/pkg/collections"
	"github.com/wcharczuk/advent/pkg/fileutil"
)

// The input
const (
	INPUT = "CRnCaCaCaSiRnBPTiMgArSiRnSiRnMgArSiRnCaFArTiTiBSiThFYCaFArCaCaSiThCaPBSiThSiThCaCaPTiRnPBSiThRnFArArCaCaSiThCaSiThSiRnMgArCaPTiBPRnFArSiThCaSiRnFArBCaSiRnCaPRnFArPMgYCaFArCaPTiTiTiBPBSiThCaPTiBPBSiRnFArBPBSiRnCaFArBPRnSiRnFArRnSiRnBFArCaFArCaCaCaSiThSiThCaCaPBPTiTiRnFArCaPTiBSiAlArPBCaCaCaCaCaSiRnMgArCaSiThFArThCaSiThCaSiRnCaFYCaSiRnFYFArFArCaSiRnFYFArCaSiRnBPMgArSiThPRnFArCaSiRnFArTiRnSiRnFYFArCaSiRnBFArCaSiRnTiMgArSiThCaSiThCaFArPRnFArSiRnFArTiTiTiTiBCaCaSiRnCaCaFYFArSiThCaPTiBPTiBCaSiThSiRnMgArCaF"
)

type xform [2]string

func main() {
	xfs := []xform{}
	fileutil.ReadByLines("../testdata/day19", func(line string) error {
		parts := strings.Split(line, " => ")
		xfs = append(xfs, xform{parts[0], parts[1]})
		return nil
	})

	results := processDistinct(xfs, INPUT)
	fmt.Printf("%d total replacements\n", results.Len())
}

func processDistinct(xf []xform, input string) *collections.SetOfString {
	ss := collections.SetOfString{}
	inputBytes := []byte(input)

	for _, xform := range xf {
		key, replacement := xform[0], xform[1]
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
	}
	return &ss
}

func slice(buffer []byte, start, end int) []byte {
	output := []byte{}
	for x := start; x < end; x++ {
		output = append(output, buffer[x])
	}
	return output
}
