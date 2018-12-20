package floats

import "math"

//IsInfinite returns true if a float64 is +Inf or -Inf
func IsInfinite(f float64) bool { return (f == math.Inf(1) || f == math.Inf(-1)) }
