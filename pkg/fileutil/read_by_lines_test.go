package fileutil

import (
	"testing"

	"github.com/blend/go-sdk/assert"
)

func TestFileReadByLines(t *testing.T) {
	assert := assert.New(t)

	tf, err := NewTempFile([]byte("foo\nbar\nbaz"))
	assert.Nil(err)
	defer tf.Close()

	called := false
	assert.Nil(ReadByLines(tf.Path, func(line string) error {
		called = true
		return nil
	}))
	assert.True(called, "We should have called the handler for `README.md`")
}
