package fileutil

import (
	"io/ioutil"
	"path/filepath"
	"regexp"

	"github.com/blend/go-sdk/exception"
)

// ListMatches lists all files in a path and filters by regex.
func ListMatches(path string, expr *regexp.Regexp) ([]string, error) {
	var files []string
	dirFiles, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, exception.New(err)
	}
	var fileBase string
	for _, dirFile := range dirFiles {
		if dirFile.IsDir() {
			continue
		}
		fileBase = filepath.Base(dirFile.Name())
		if expr == nil {
			files = append(files, fileBase)
			continue
		}
		if expr.MatchString(fileBase) {
			files = append(files, fileBase)
		}
	}
	return files, nil
}
