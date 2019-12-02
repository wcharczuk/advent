package main

import (
	"testing"

	"github.com/blend/go-sdk/assert"
)

func Test_Intcode(t *testing.T) {
	assert := assert.New(t)

	testCases := [...]struct {
		Input    []int
		Expected []int
	}{
		{Input: []int{1, 0, 0, 0, 99}, Expected: []int{2, 0, 0, 0, 99}},
		{Input: []int{2, 3, 0, 3, 99}, Expected: []int{2, 3, 0, 6, 99}},
		{Input: []int{2, 4, 4, 5, 99, 0}, Expected: []int{2, 4, 4, 5, 99, 9801}},
		{Input: []int{2, 4, 4, 5, 99, 0}, Expected: []int{2, 4, 4, 5, 99, 9801}},
		{Input: []int{1, 1, 1, 4, 99, 5, 6, 0, 99}, Expected: []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
		{Input: []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}, Expected: []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}},
	}

	for _, tc := range testCases {
		result, err := Intcode(tc.Input...)
		assert.Nil(err)
		assert.Equal(tc.Expected, result)
	}
}
