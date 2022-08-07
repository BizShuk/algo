package gomock

type useLessObject struct {
}

// [Pattern]: [Go Testing] Suck as function should be immutable (no side effect), just call it and test return value
func UseLessObject() useLessObject {
	return useLessObject{}
}
