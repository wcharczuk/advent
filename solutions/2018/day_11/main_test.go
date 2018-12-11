package main

import (
	"testing"

	"github.com/blend/go-sdk/assert"
)

func TestPowerLevel(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(-5, powerLevel(57, 122, 79))
	assert.Equal(0, powerLevel(39, 217, 196))
	assert.Equal(4, powerLevel(71, 101, 153))
}
