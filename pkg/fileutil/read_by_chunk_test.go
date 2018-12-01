package fileutil

import (
	"testing"

	"github.com/blend/go-sdk/assert"
)

func TestFileReadByChunks(t *testing.T) {
	assert := assert.New(t)

	called := false
	assert.Nil(ReadByChunks("README.md", 32, func(chunk []byte) error {
		called = true
		return nil
	}))
	assert.True(called, "We should have called the handler for `README.md`")
}
