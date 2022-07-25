package array

// Notice: bucket sorting without heap, https://leetcode.com/problems/top-k-frequent-elements/submissions/

func TopKElement(nums []int, k int) []int {
	freqCounter := make(map[int]int)

	for _, v := range nums {
		freqCounter[v] += 1
	}

	bucket := make([][]int, len(nums)+1)
	for k, v := range freqCounter {
		bucket[v] = append(bucket[v], k)
	}
	result := []int{}
	for i := len(bucket) - 1; i >= 0; i-- {
		for j := 0; j < len(bucket[i]); j++ { // this filter out the bucket don't have anything
			result = append(result, bucket[i][j])
			if len(result) == k {
				return result
			}
		}
	}
	return []int{}
}
