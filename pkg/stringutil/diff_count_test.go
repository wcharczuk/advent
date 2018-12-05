package stringutil

import (
	"testing"

	"github.com/blend/go-sdk/assert"
)

type diffCountTestCase struct {
	A, B     string
	Expected int
}

func TestDiffCount(t *testing.T) {
	it := assert.New(t)

	testCases := []diffCountTestCase{
		{A: "", B: "", Expected: 0},
		{A: "aaa", B: "", Expected: -1},
		{A: "", B: "aaa", Expected: -1},
		{A: "aaa", B: "aaa", Expected: 0},
		{A: "aba", B: "aaa", Expected: 1},
		{A: "aaa", B: "aba", Expected: 1},
		{A: "aac", B: "aba", Expected: 2},
		{A: "dac", B: "aba", Expected: 3},
	}

	for _, tc := range testCases {
		it.Equal(tc.Expected, DiffCount(tc.A, tc.B))
	}
}
