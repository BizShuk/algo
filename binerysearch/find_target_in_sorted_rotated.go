package binerysearch

// Notice: Don't try to optimize the multi-dimention if-else condition, just LIST IT OUTS

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		m := (l + r) / 2
		mNum := nums[m]

		if mNum == target {
			return m
		}

		if nums[m] > nums[r] { // m > l > r, big/small are at right
			if target > nums[r] && target < nums[m] {
				r = m - 1
			} else {
				l = m + 1
			}
		} else if nums[l] > nums[m] { // l > r > m, big/small are at left
			if target > nums[m] && target < nums[l] {
				l = m + 1
			} else {
				r = m - 1
			}
		} else { // standard bst
			if target > mNum {
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}

	return -1
}
