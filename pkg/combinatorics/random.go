package combinatorics

import (
	"math/rand"
	"time"
)

var (
	provider = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// RandomInt returns a random int from an array.
func RandomInt(values ...int) int {
	if len(values) == 0 {
		return 0
	}
	if len(values) == 1 {
		return values[0]
	}
	return values[provider.Intn(len(values))]
}

// RandomFloat64 returns a random int from an array.
func RandomFloat64(values ...float64) float64 {
	if len(values) == 0 {
		return 0
	}
	if len(values) == 1 {
		return values[0]
	}
	return values[provider.Intn(len(values))]
}

// RandomString returns a random string from an array.
func RandomString(values ...string) string {
	if len(values) == 0 {
		return ""
	}
	if len(values) == 1 {
		return values[0]
	}
	return values[provider.Intn(len(values))]
}
