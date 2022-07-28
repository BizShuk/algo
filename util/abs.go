package util

// [Pattern]: Abs
func Abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}
