package util

// [Pattern]: [Go Util] Abs
func Abs[K Number](x, y K) K {
	if x > y {
		return x - y
	}
	return y - x
}
