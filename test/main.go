package main

import "fmt"

func main() {
	s := []int{
		100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 99, 97,
	}
	fmt.Println(canPartition(s))
	// s2 := []int{
	// 100, 50, 50,
	// }
	// fmt.Println(canPartition(s2))
}

// 1. if num is odd => can't be partitioned
// 2. the partition should be equal to (total sum / 2)
//    => now we have target sum
//    => N nums sum equal to target sum
func canPartition(nums []int) bool {
	totalSum := 0
	max := 0
	for _, num := range nums {
		totalSum += num
		max = Max(max, num)
	}

	target := totalSum / 2

	if totalSum%2 == 1 || max > target {
		return false
	}

	if max == target {
		return true
	}

	dp := map[int]bool{}

	// DP solution:
	dp[0] = true
	count := 0
	for i := 0; i < len(nums); i++ {
		nextDp := map[int]bool{}
		count += len(dp)
		for k := range dp {
			if k+nums[i] == target {
				return true
			}
			nextDp[k+nums[i]] = true
			nextDp[k] = true
		}
		dp = nextDp
	}
	fmt.Println(count)
	return false

	//
	// var dfs func(nums []int, i int, sum int) bool
	// count := 0
	// dfs = func(nums []int, i int, sum int) bool {
	// 	if i >= len(nums) {
	// 		count += 1
	// 		fmt.Println(count)
	// 		return false
	// 	}

	// 	if sum == target || sum+nums[i] == target {
	// 		return true
	// 	}
	// 	count += 1
	// 	fmt.Println(count)
	// 	return dfs(nums, i+1, sum+nums[i]) || dfs(nums, i+1, sum)
	// }

	// return dfs(nums, 0, 0)
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
