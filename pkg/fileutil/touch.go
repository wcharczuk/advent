package fileutil

import (
	"os"
)

// Touch creates and closes a file.
func Touch(filePath string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}
