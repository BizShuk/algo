package backtrack

import "sort"

// [Variant]: [Combination BT] Combination Sum with unlimited selection
func CombinationSum_dfs(candidates []int, target int) [][]int {
	result := [][]int{}

	var dfs func(subset []int, i int, sum int)

	dfs = func(subset []int, i int, sum int) {

		if sum > target {
			return
		}

		if sum == target {
			result = append(result, append([]int{}, subset...))
			return
		}

		// [Way 1]:
		for j, candidate := range candidates {
			if j < i {
				continue
			}
			dfs(append(subset, candidate), j, sum+candidate)
		}

		// [Way 2]: Using i ask current pick-up
		// if i >= len(candidates) {
		// 	return
		// }

		// dfs(subset, i+1, sum) // 不取 往下一個走
		// subset = append(subset, candidates[i])
		// dfs(subset, i, sum+candidates[i])
	}

	dfs([]int{}, 0, 0)
	return result
}

// [Variant]: [Combination BT] Combination sum with unique subset
func CombinationSum2(candidates []int, target int) [][]int {
	result := [][]int{}

	sort.Ints(candidates)

	// [Way 1]:
	var dfs func(subset []int, i int, sum int)

	dfs = func(subset []int, i int, sum int) {

		if sum > target || i >= len(candidates) {
			return
		}

		if sum+candidates[i] == target {
			subset = append(subset, candidates[i])
			result = append(result, subset)
			return
		}

		subset = append(subset, candidates[i])
		dfs(subset, i+1, sum+candidates[i])
		subset = subset[: len(subset)-1 : len(subset)-1]

		j := i // Find next num index
		for j < len(candidates) && candidates[j] <= candidates[i] {
			j += 1
		}
		dfs(subset, j, sum)
	}

	// [Way 2]: on python
	// def backtrack(cur, pos, target):
	// if target == 0:
	// 		res.append(cur.copy())
	// if target <= 0:
	// 		return

	// prev = -1
	// for i in range(pos, len(candidates)):
	// 		if candidates[i] == prev:
	// 			continue
	// 		cur.append(candidates[i])
	// 		backtrack(cur, i + 1, target - candidates[i])
	// 		cur.pop()
	// 		prev = candidates[i]

	dfs([]int{}, 0, 0)

	return result
}
