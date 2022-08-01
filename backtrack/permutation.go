package backtrack

// [Variant]: [Combination BT] Permutation, using map as filter instead of index
func Permutation(nums []int) [][]int {
	result := [][]int{}

	var dfs func(per []int, set map[int]bool, count int)

	dfs = func(per []int, set map[int]bool, count int) {
		if count == len(nums) {
			result = append(result, per)
		}

		for _, num := range nums {
			if set[num] {
				continue
			}
			set[num] = true
			dfs(append(per, num), set, count+1)
			delete(set, num)
		}

	}
	dfs([]int{}, map[int]bool{}, 0)
	return result

}
