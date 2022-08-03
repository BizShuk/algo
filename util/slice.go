package util

func Pop[T1 any](s []T1) []T1 {
	return s[: len(s)-1 : len(s)-1]
}

func Filter[T1 any](s []T1, f func(T1) bool) (output []T1) {
	for _, v := range s {
		if f(v) {
			output = append(output, v)
		}
	}
	return output
}

func Map[T1 any, T2 any](s []T1, f func(T1) T2) (output []T2) {
	for _, v := range s {
		output = append(output, f(v))
	}
	return output
}

func ReduceWithInitialValue[T1, T2 any](input []T1, acc T2, f func(T2, T1) T2) T2 {
	for _, v := range input {
		acc = f(acc, v)
	}

	return acc
}
