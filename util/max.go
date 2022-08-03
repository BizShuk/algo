package util

// [Pattern]: Max
func Max[K Number](x, y K) K {
	if x > y {
		return x
	}
	return y
}

func MaxNegotive[K Number](x, y K) K {
	if x > 0 {
		return y
	}

	if x > y {
		return x
	}

	return y
}
