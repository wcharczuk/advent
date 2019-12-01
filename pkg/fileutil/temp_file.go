package fileutil

import (
	"io/ioutil"
	"os"
)

// NewTempFile creates a new temporary file that deletes on close.
func NewTempFile(contents []byte) (*TempFile, error) {
	tf, err := ioutil.TempFile("", "")
	if err != nil {
		return nil, err
	}
	defer tf.Close()
	_, err = tf.Write(contents)
	if err != nil {
		return nil, err
	}
	return &TempFile{
		Path: tf.Name(),
	}, nil
}

// TempFile is a file reference at a given path
// that removes itself on close.
type TempFile struct {
	Path string
}

// Close deletes the file.
func (tf *TempFile) Close() error {
	return os.Remove(tf.Path)
}
