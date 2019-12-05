package main

import (
	"testing"

	"github.com/blend/go-sdk/assert"
)

func Test_ParseOpcode(t *testing.T) {
	assert := assert.New(t)

	parsed := ParseOpcode(1002)
	assert.Equal(2, parsed.Op)
	assert.Equal(0, parsed.Modes[0])
	assert.Equal(1, parsed.Modes[1])
	assert.Equal(0, parsed.Modes[2])
}
