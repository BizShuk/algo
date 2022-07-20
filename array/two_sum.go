package array

func twoSum(s []int, target int) bool {
	x := make(map[int]bool)

	for _, e := range s {
		if x[target-e] {
			return true
		}
		x[e] = true
	}
	return false
}
