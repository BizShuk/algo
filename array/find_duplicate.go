package array

// [Notice]: Leverage index to find the slot for each [1,n] value
// N+1 elements with [1,n] value => one element must be duplicate
func FindDuplicate(nums []int) int {
	numMap := map[int]bool{}

	for _, num := range nums {
		if numMap[num] {
			return num
		}
		numMap[num] = true
	}
	return -1
}
