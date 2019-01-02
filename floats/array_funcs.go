package floats

//AllEqual takes an array of floats and a tolerance for comparison, and returns true if all members of the array are equal under that tolerance.
func AllEqual(floats []float64, tolerance float64) bool {
	ref := floats[0]
	for _, f := range floats {
		deviation := f / ref
		if deviation < 1.0-tolerance || deviation > 1.0+tolerance {
			return false
		}
	}
	return true
}

//All is a higher order function that takes an array and a predicate, and returns true if all members of the array satisfy the prediacte
func All(vs []float64, f func(float64) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

//Any is a higher order function that takes an array and a predicate, and returns true if any member of the array satisfies the predicate
func Any(vs []float64, f func(float64) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}
