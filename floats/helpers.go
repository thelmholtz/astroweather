package floats

import "math"

//IsInfinite returns true if a float64 is +Inf or -Inf
func IsInfinite(f float64) bool { return (f == math.Inf(1) || f == math.Inf(-1)) }

//Equals returns if two floats are equal at the given tolerance
func Equals(f1, f2, tolerance float64) bool {
	deviation := math.Abs(f1 / f2)
	if deviation < 1.0-tolerance || deviation > 1.0+tolerance {
		return false
	}
	return true
}
