package util

// [Pattern]: Min
func Min[K Number](x, y K) K {
	if x < y {
		return x
	}
	return y
}

// x might be initial value(invalid)
func MinPositive[K Number](x, y K) K {
	if x < 0 {
		return y
	}
	if x < y {
		return x
	}
	return y
}
