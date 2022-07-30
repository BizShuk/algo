package util

// [Pattern]: Max
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func MaxNegotive(x, y int) int {
	if x > 0 {
		return y
	}

	if x > y {
		return x
	}

	return y
}
