package fileutil

import (
	"os"

	"github.com/blend/go-sdk/exception"
)

// Touch creates and closes a file.
func Touch(filePath string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return exception.New(err)
	}
	defer f.Close()
	return nil
}
