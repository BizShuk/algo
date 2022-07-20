package DSmap

// by element existing
func Duplicate(x int, set map[int]bool) bool {
	if set[x] {
		return true
	}
	return false
}

// by map size
func DuplicateBySize(set map[int]bool, list []int) bool {
	if len(set) < len(list) {
		return true
	}
	return false
}
