package fileutil

import (
	"os"
	"path/filepath"
)

// ListRecursive lists a path recusively.
func ListRecursive(path string) ([]os.FileInfo, error) {
	return ListRecursiveMatches(path, nil)
}

// ListRecursiveMatches lists all files and directories recusively under a folder that optionally match a predicate.
// The predicate should return a bool if the fileinfo should be included in the results.
// It should return `filepath.SkipDir` if you don't want to process a given directory.
func ListRecursiveMatches(path string, predicate func(os.FileInfo) (bool, error)) (output []os.FileInfo, err error) {
	var shouldAdd bool
	var skipDir error
	err = filepath.Walk(path, func(currentPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if path == currentPath {
			return nil
		}

		if predicate == nil {
			output = append(output, info)
		}

		shouldAdd, skipDir = predicate(info)
		if shouldAdd {
			output = append(output, info)
		}
		return skipDir
	})
	return
}
