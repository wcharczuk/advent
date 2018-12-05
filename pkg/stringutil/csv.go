package stringutil

import "strings"

// CSV returns a comma separated string of a given slice of values.
func CSV(values []string) string {
	return strings.Join(values, ",")
}

// SplitCSV returns a comma separated string of a given slice of values.
func SplitCSV(csv string) (output []string) {
	pieces := strings.Split(csv, ",")
	for _, str := range pieces {
		if str != "" {
			output = append(output, strings.TrimSpace(str))
		}
	}
	return
}
