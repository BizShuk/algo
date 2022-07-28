package twopointers

import "sort"

// Notice: condition boundary and using continue/break if the flow should not execute any more
// For i, i+= 1 and move forward if duplicate
// For lp, it just need to handle if sum == 0 for duplicate case
func ThreeSum(nums []int) [][]int {
	result := [][]int{}
	sort.Ints(nums)

	i := 0
	for i < len(nums)-2 {
		if i > 0 && nums[i] == nums[i-1] {
			i += 1
			continue
		}
		fixNum := nums[i]

		lp, rp := i+1, len(nums)-1

		for lp < rp {

			sum := fixNum + nums[lp] + nums[rp]

			if sum > 0 {
				rp -= 1
				continue
			}
			if sum < 0 {
				lp += 1
				continue
			}

			if sum == 0 {
				result = append(result, []int{fixNum, nums[lp], nums[rp]})
				lp += 1
				for nums[lp] == nums[lp-1] && lp < rp {
					lp += 1
				}
			}

		}
		i += 1
	}

	return result
}
