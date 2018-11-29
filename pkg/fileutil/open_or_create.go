package fileutil

import (
	"os"

	"github.com/blend/go-sdk/exception"
)

// OpenOrCreate opens or creates a file.
func OpenOrCreate(filePath string) (*os.File, error) {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if os.IsNotExist(err) {
		f, err = os.Create(filePath)
		return f, exception.New(err)
	}
	return f, exception.New(err)
}
