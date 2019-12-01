package fileutil

import "strconv"

// ReadFloats reads a file and parses each line as a float64.
func ReadFloats(file string) (values []float64, err error) {
	err = ReadByLines(file, func(line string) error {
		value, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return err
		}
		values = append(values, value)
		return nil
	})
	return
}
