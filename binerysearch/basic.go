package binerysearch

// Notice: [Pattern] two pointers basic, indexs/2
// 1. when moving left pointer and right pointer, make sure add/minus 1
// 2. (left pointer + right pointer)/2 if the sum is odd, result will be the left pointer(less)

// Notice: Binary Search: for l "<=" r,

// return index if found, otherwise return -1
func findTarget(s []int, target int) int {
	l, r := 0, len(s)-1

	for l <= r {
		m := (l + r) / 2
		if s[m] == target { // do something or return
			return m // might continue execute instead break/return
		}

		// move l/r, Be aware of equal condition if target matched isn't return/break
		if target > s[m] {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return -1
}
