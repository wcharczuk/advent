package fileutil

import (
	"bufio"
	"os"
)

//ReadLineHandler is a receiver for a line of a file.
type ReadLineHandler func(line string) error

// ReadByLines reads a file and calls the handler for each line.
func ReadByLines(filePath string, handler ReadLineHandler) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		err = handler(line)
		if err != nil {
			return err
		}
	}

	return nil
}
