package seq

import (
	"testing"

	"github.com/blend/go-sdk/assert"
)

func TestInt(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{0, 1, 2, 3}, Int(3).Values())
	assert.Equal([]int{1, 2, 3}, Int(3).WithStart(1).Values())
	assert.Empty(Int(3).WithStart(3).Values())
}

func TestIntPanic(t *testing.T) {
	assert := assert.New(t)

	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()

		(&IntSeq{}).Values()
	}()
	assert.True(didPanic)
}
