package fileutil

import (
	"io/ioutil"
	"os"
)

// List returns all the file infos for a given directory path.
func List(path string) ([]os.FileInfo, error) {
	return ListMatches(path, nil)
}

// ListMatches returns all files in a path filtered by an optional predicate.
func ListMatches(path string, predicate func(os.FileInfo) bool) ([]os.FileInfo, error) {
	dirFiles, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var files []os.FileInfo
	for _, dirFile := range dirFiles {
		if dirFile.IsDir() {
			continue
		}
		if predicate == nil || (predicate != nil && predicate(dirFile)) {
			files = append(files, dirFile)
		}
	}
	return files, nil
}
