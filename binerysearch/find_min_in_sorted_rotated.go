package binerysearch

// Notice: If Condition logic
// Assume finding minimum by going left (rp = mp - 1)
// address exception case first which left rotated

func findMin(nums []int) int {
	lp, rp := 0, len(nums)-1

	min := 5001

	for lp <= rp {
		mp := (lp + rp) / 2

		if nums[mp] < min {
			min = nums[mp]
		}

		if nums[mp] >= nums[lp] && nums[mp] > nums[rp] {
			lp = mp + 1
		} else {
			rp = mp - 1
		}
	}
	return min
}
