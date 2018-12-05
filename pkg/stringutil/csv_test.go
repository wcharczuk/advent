package stringutil

import (
	"testing"

	"github.com/blend/go-sdk/assert"
)

type splitCSVTestCase struct {
	Input    string
	Expected []string
}

func TestSplitCSV(t *testing.T) {
	assert := assert.New(t)

	testCases := []splitCSVTestCase{
		{},
		{Input: ","},
		{Input: "a,", Expected: []string{"a"}},
		{Input: "a,b,c", Expected: []string{"a", "b", "c"}},
		{Input: "a,b,c,d", Expected: []string{"a", "b", "c", "d"}},
		{Input: "a, b, c", Expected: []string{"a", "b", "c"}},
		{Input: "a,\tb,\tc", Expected: []string{"a", "b", "c"}},
	}

	for _, tc := range testCases {
		assert.Equal(tc.Expected, SplitCSV(tc.Input), tc.Input)
	}
}
