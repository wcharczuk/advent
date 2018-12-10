package main

import (
	"testing"

	"github.com/blend/go-sdk/assert"
)

func TestRingAdd(t *testing.T) {
	assert := assert.New(t)

	r := NewRing()

	// [-] (0)
	r = r.Add(0)
	assert.Equal(0, r.Value)
	assert.NotNil(r.CCW)
	assert.NotNil(0, r.CCW.Value)
	assert.NotNil(r.CW)
	assert.Equal(0, r.CW.Value)
	assert.Equal("(0)", r.String())

	// [1]  0 (1)
	r = r.Add(1)
	assert.Equal(1, r.Value)
	assert.NotNil(r.CCW)
	assert.Equal(0, r.CCW.Value)
	assert.NotNil(r.CW)
	assert.Equal(0, r.CW.Value)
	assert.Equal("0 (1)", r.String())

	// [2]  0 (2) 1
	r = r.Add(2)
	assert.Equal(2, r.Value)
	assert.NotNil(r.CCW)
	assert.Equal(0, r.CCW.Value)
	assert.NotNil(r.CW)
	assert.Equal(1, r.CW.Value)
	assert.Equal("0 (2) 1", r.String())

	// [3]  0  2  1 (3)
	r = r.Add(3)
	assert.Equal(3, r.Value)
	assert.NotNil(r.CCW)
	assert.Equal(1, r.CCW.Value)
	assert.NotNil(r.CW)
	assert.Equal(0, r.CW.Value)

	// [4]  0 (4) 2  1  3
	r = r.Add(4)
	assert.Equal(4, r.Value)
	assert.NotNil(r.CW)
	assert.Equal(2, r.CW.Value)
	assert.NotNil(r.CCW)
	assert.Equal(0, r.CCW.Value)

	// [5]  0  4  2 (5) 1  3
	r = r.Add(5)
	assert.Equal(5, r.Value)
	assert.NotNil(r.CCW)
	assert.Equal(2, r.CCW.Value)
	assert.NotNil(r.CW)
	assert.Equal(1, r.CW.Value)

	// [6]  0  4  2  5  1 (6) 3
	r = r.Add(6)
	assert.Equal(6, r.Value)
	assert.NotNil(r.CCW)
	assert.Equal(1, r.CCW.Value)
	assert.NotNil(r.CW)
	assert.Equal(3, r.CW.Value)
	assert.Equal("0 4 2 5 1 (6) 3", r.String())

	// [7]  0  4  2  5  1  6  3 (7)
	r = r.Add(7)
	assert.Equal(7, r.Value)
	assert.NotNil(r.CCW)
	assert.Equal(3, r.CCW.Value)
	assert.NotNil(r.CW)
	assert.Equal(0, r.CW.Value)
	assert.Equal("0 4 2 5 1 6 3 (7)", r.String())
}
