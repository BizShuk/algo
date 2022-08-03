package util

import "golang.org/x/exp/constraints"

func MapKeys[T1 constraints.Ordered, T2 any](s map[T1]T2) (output []T1) {
	for k := range s {
		output = append(output, k)
	}
	return output
}

func MapValues[T1 constraints.Ordered, T2 any](s map[T1]T2) (output []T2) {
	for _, v := range s {
		output = append(output, v)
	}
	return output
}
