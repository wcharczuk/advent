package stringutil

import (
	"fmt"
	"testing"

	"github.com/blend/go-sdk/assert"
)

func TestCaselessDiffCount(t *testing.T) {
	it := assert.New(t)

	testCases := []diffCountTestCase{
		{A: "", B: "", Expected: 0},
		{A: "aaa", B: "", Expected: -1},
		{A: "", B: "aaa", Expected: -1},
		{A: "aaa", B: "aaa", Expected: 0},
		{A: "aaa", B: "aAa", Expected: 0},
		{A: "aba", B: "aaa", Expected: 1},
		{A: "aaa", B: "aba", Expected: 1},
		{A: "aac", B: "aba", Expected: 2},
		{A: "dac", B: "aba", Expected: 3},
	}

	for _, tc := range testCases {
		it.Equal(tc.Expected, CaselessDiffCount(tc.A, tc.B), fmt.Sprintf("%s vs. %s", tc.A, tc.B))
	}
}
