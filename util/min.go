package util

// [Pattern]: Min
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// x might be initial value(invalid)
func MinPositive(x, y int) int {
	if x < 0 {
		return y
	}
	if x < y {
		return x
	}
	return y
}
