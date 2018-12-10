package main

import (
	"testing"

	"github.com/blend/go-sdk/assert"
)

func TestRingAdd(t *testing.T) {
	assert := assert.New(t)

	r := NewRing()

	// [-] (0)
	r.Add(0)
	assert.Equal(0, r.Cursor)
	assert.Len(r.Values, 1)
	assert.Equal(0, r.Values[0])

	// [1]  0 (1)
	r.Add(1)
	assert.Equal(1, r.Cursor)
	assert.Len(r.Values, 2)
	assert.Equal(0, r.Values[0])
	assert.Equal(1, r.Values[1])

	// [2]  0 (2) 1
	r.Add(2)
	assert.Equal(1, r.Cursor)
	assert.Len(r.Values, 3)
	assert.Equal(0, r.Values[0])
	assert.Equal(2, r.Values[1])
	assert.Equal(1, r.Values[2])

	// [3]  0  2  1 (3)
	r.Add(3)
	assert.Equal(3, r.Cursor)
	assert.Len(r.Values, 4)
	assert.Equal(0, r.Values[0])
	assert.Equal(2, r.Values[1])
	assert.Equal(1, r.Values[2])
	assert.Equal(3, r.Values[3])

	// [4]  0 (4) 2  1  3
	r.Add(4)
	assert.Equal(1, r.Cursor, r.String())
	assert.Len(r.Values, 5)
	assert.Equal(0, r.Values[0])
	assert.Equal(4, r.Values[1])
	assert.Equal(2, r.Values[2])
	assert.Equal(1, r.Values[3])
	assert.Equal(3, r.Values[4])

	// [5]  0  4  2 (5) 1  3
	r.Add(5)
	assert.Equal(3, r.Cursor, r.String())
	assert.Len(r.Values, 6)
	assert.Equal(0, r.Values[0])
	assert.Equal(4, r.Values[1])
	assert.Equal(2, r.Values[2])
	assert.Equal(5, r.Values[3])
	assert.Equal(1, r.Values[4])
	assert.Equal(3, r.Values[5])

	// [6]  0  4  2  5  1 (6) 3
	r.Add(6)
	assert.Equal(5, r.Cursor, r.String())
	assert.Len(r.Values, 7)
	assert.Equal(0, r.Values[0])
	assert.Equal(4, r.Values[1])
	assert.Equal(2, r.Values[2])
	assert.Equal(5, r.Values[3])
	assert.Equal(1, r.Values[4])
	assert.Equal(6, r.Values[5])
	assert.Equal(3, r.Values[6])

	// [7]  0  4  2  5  1  6  3 (7)
	r.Add(7)
	assert.Equal(7, r.Cursor, r.String())
	assert.Len(r.Values, 8)
	assert.Equal(0, r.Values[0])
	assert.Equal(4, r.Values[1])
	assert.Equal(2, r.Values[2])
	assert.Equal(5, r.Values[3])
	assert.Equal(1, r.Values[4])
	assert.Equal(6, r.Values[5])
	assert.Equal(3, r.Values[6])
	assert.Equal(7, r.Values[7])
}

func TestRingRemoveFromCursor(t *testing.T) {
	assert := assert.New(t)

	r := NewRing()
	for x := 0; x < 23; x++ {
		r.Add(x)
	}
	assert.Equal(9, r.Remove())
	assert.Equal(6, r.Cursor, r.String())
	assert.Len(r.Values, 22)
}

func TestRemoveAt(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{2, 3}, removeAt([]int{1, 2, 3}, 0))
	assert.Equal([]int{1, 3}, removeAt([]int{1, 2, 3}, 1))
	assert.Equal([]int{1, 2}, removeAt([]int{1, 2, 3}, 2))

	assert.Equal([]int{2, 3, 4}, removeAt([]int{1, 2, 3, 4}, 0))
	assert.Equal([]int{1, 3, 4}, removeAt([]int{1, 2, 3, 4}, 1))
	assert.Equal([]int{1, 2, 4}, removeAt([]int{1, 2, 3, 4}, 2))
	assert.Equal([]int{1, 2, 3}, removeAt([]int{1, 2, 3, 4}, 3))
}

func TestCCW(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(1, ccw(2, 3, 1))
	assert.Equal(0, ccw(2, 3, 2))
	assert.Equal(2, ccw(2, 3, 3))
	assert.Equal(2, ccw(2, 3, 6))
}
