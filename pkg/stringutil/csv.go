package stringutil

import "strings"

// CSV returns a comma separated string of a given slice of values.
func CSV(values []string) string {
	return strings.Join(values, ",")
}
