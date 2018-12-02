package fileutil

import (
	"testing"

	"github.com/blend/go-sdk/assert"
)

func TestFileReadByChunks(t *testing.T) {
	assert := assert.New(t)

	tf, err := NewTempFile([]byte("abcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabc"))
	assert.Nil(err)
	defer tf.Close()

	called := false
	assert.Nil(ReadByChunks(tf.Path, 32, func(chunk []byte) error {
		called = true
		return nil
	}))
	assert.True(called, "We should have called the handler for `README.md`")
}
