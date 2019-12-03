package main

import (
	"testing"

	"github.com/blend/assert/assert"
)

func Test_ParsePath(t *testing.T) {
	assert := assert.New(t)

	path := "R1009,U993,L383,D725,R163"
	parsed := ParsePath(path)
	assert.Len(parsed, 5)

	assert.Equal("R", parsed[0].Direction)
	assert.Equal(1009, parsed[0].Distance)

	assert.Equal("U", parsed[1].Direction)
	assert.Equal(993, parsed[1].Distance)

	assert.Equal("L", parsed[2].Direction)
	assert.Equal(383, parsed[2].Distance)

	assert.Equal("D", parsed[3].Direction)
	assert.Equal(725, parsed[3].Distance)

	assert.Equal("R", parsed[4].Direction)
	assert.Equal(163, parsed[4].Distance)
}

func Test_Point_Distance(t *testing.T) {
	assert := assert.New(t)

	testCases := [...]struct {
		Input    Point
		Expected int
	}{
		{Input: Point{1, 1}, Expected: 2},
		{Input: Point{-1, 1}, Expected: 2},
		{Input: Point{1, -1}, Expected: 2},
		{Input: Point{-1, -1}, Expected: 2},
	}

	for _, tc := range testCases {
		assert.Equal(tc.Expected, tc.Input.Distance())
	}
}

func Test_Path_Expand(t *testing.T) {
	assert := assert.New(t)

	path := Path([]Segment{
		{
			Direction: "R",
			Distance:  10,
		},
		{
			Direction: "U",
			Distance:  9,
		},
		{
			Direction: "L",
			Distance:  8,
		},
		{
			Direction: "D",
			Distance:  7,
		},
	})

	points := path.Expand()
	assert.Len(points, 34)
	assert.Equal("1,0", points[0].String())
	assert.Equal("10,0", points[9].String())
	assert.Equal("10,9", points[18].String())
	assert.Equal("2,9", points[26].String())
	assert.Equal("2,2", points[33].String())
}
