package backtrack

// [Variant]: [Combination BT] Permutation, using map as filter instead of index
func Permutation(nums []int) [][]int {

	result := [][]int{}
	set := map[int]bool{}
	var dfs func(per []int)

	dfs = func(per []int) {
		if len(per) == len(nums) {
			result = append(result, per)
			return
		}

		for _, num := range nums {
			if set[num] {
				continue
			}
			set[num] = true
			dfs(append(per, num))
			delete(set, num)
		}

	}
	dfs([]int{})
	return result
}
