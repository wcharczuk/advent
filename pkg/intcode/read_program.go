package intcode

import (
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// ReadProgramFile reads a program from a file.
func ReadProgramFile(filePath string) ([]int, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ReadProgram(f)
}

// ReadProgram reads a program as a csv of ints.
func ReadProgram(r io.Reader) ([]int, error) {
	contents, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	rawValues := strings.Split(string(contents), ",")
	opCodes := make([]int, len(rawValues))
	for x := 0; x < len(rawValues); x++ {
		opCodes[x], err = strconv.Atoi(strings.TrimSpace(rawValues[x]))
		if err != nil {
			return nil, err
		}
	}
	return opCodes, nil
}
