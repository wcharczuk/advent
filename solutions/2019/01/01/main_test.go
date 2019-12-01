package main

import (
	"testing"

	"github.com/blend/go-sdk/assert"
)

func Test_ComputeFuel(t *testing.T) {
	assert := assert.New(t)

	tcs := [...]struct {
		Input    float64
		Expected float64
	}{
		{Input: 12, Expected: 2},
		{Input: 14, Expected: 2},
		{Input: 1969, Expected: 654},
		{Input: 100756, Expected: 33583},
	}

	for _, tc := range tcs {
		assert.Equal(tc.Expected, ComputeFuel(tc.Input))
	}
}
