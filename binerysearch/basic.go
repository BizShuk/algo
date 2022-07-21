package binerysearch

// Notice: two pointers basic, indexs/2
// 1. when moving left pointer and right pointer, make sure add/minus 1
// 2. (left pointer + right pointer)/2 if the sum is odd, result will be the left pointer(less)

// return index if found, otherwise return -1
func findTargetInSortedSlice(s []int, target int) int {
	leftP, rightP := 0, len(s)-1

	for leftP <= rightP {
		midP := (leftP + rightP) / 2
		if s[midP] == target {
			return midP
		}

		if target > s[midP] {
			leftP = midP + 1
		} else {
			rightP = midP - 1
		}
	}

	return -1
}
