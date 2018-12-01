package fileutil

import (
	"testing"

	"github.com/blend/go-sdk/assert"
)

func TestFileReadByLines(t *testing.T) {
	assert := assert.New(t)

	called := false
	assert.Nil(ReadByLines("README.md", func(line string) error {
		called = true
		return nil
	}))
	assert.True(called, "We should have called the handler for `README.md`")
}
