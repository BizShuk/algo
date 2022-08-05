package backtrack

import "sort"

// [Pattern]: [Combination BT] Generate all unique subsets
func Subsets(nums []int) [][]int {
	result := [][]int{}

	var bt func(nums []int, subset []int, i int)
	bt = func(nums []int, subset []int, i int) {

		if i < 0 {
			x := make([]int, len(subset))
			copy(x, subset)
			// Warning: sorting will mess up the value with index.
			// TODO: sorting cause value mismatch in subset of parent function
			// find out problem exactly with cap and len on slice
			sort.Ints(x)
			result = append(result, x)
			return
		}

		bt(nums, subset, i-1)
		subset = append(subset, nums[i])
		bt(nums, subset, i-1)
	}

	bt(nums, []int{}, len(nums)-1)
	return result
}

// [Variant]: [Combination BT] Generate all subset with duplicate number
func SubsetsWithDup(nums []int) [][]int {
	result := [][]int{
		{},
	}

	// [Way 1]:
	counts := map[int]int{}

	for _, num := range nums {
		counts[num] += 1
	}

	for num, count := range counts {
		appendSet := []int{num}
		endIndex := len(result)
		for i := 0; i < count; i++ {
			t := appendSet
			for j := 0; j < endIndex; j++ {
				x := append(result[j], t...)
				result = append(result, x)
			}
			appendSet = append(appendSet, num)
		}
	}

	// [Way 2]:
	//   def backtrack(i, subset):
	//         if i == len(nums):
	//             res.append(subset[::])
	//             return

	//         # All subsets that include nums[i]
	//         subset.append(nums[i])
	//         backtrack(i + 1, subset)
	//         subset.pop()
	//         # All subsets that don't include nums[i]
	//         while i + 1 < len(nums) and nums[i] == nums[i + 1]:
	//             i += 1
	//         backtrack(i + 1, subset)

	return result
}
