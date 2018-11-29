package mathutil

import "math"

// PowOfInt returns the base to the power.
func PowOfInt(base, power uint) int {
	if base == 2 {
		return 1 << power
	}
	return int(math.RoundToEven(math.Pow(float64(base), float64(power))))
}
