package dp

import "github.com/bizshuk/algo/util"

// [Notice]: Take anything before it to calculate
// [Pattern]: [Simple DP] RobHouse
// Climbing Stair/Fibonacci are in this category
func Rob(nums []int) int {
	untake, take := 0, 0

	for _, num := range nums {
		tmp := untake + num
		untake = util.Max(untake, take)
		take = tmp
	}

	return util.Max(untake, take)
}

func Rob2(nums []int) int {
	dp := [][]int{
		{0, 0},
	}
	max := 0

	for i, num := range nums {
		cdp := make([]int, 2)
		cdp[0] = util.Max(dp[i][0], dp[i][1])
		cdp[1] = dp[i][0] + num
		dp = append(dp, cdp)
		max = util.Max(util.Max(cdp[0], cdp[1]), max)
	}

	return max
}

// [Variant]: [Simple DP] RobHouse with cycle
// Assume two expressions for non-taking first or last elements
func RobWithCycle(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return util.Max(Rob(nums[1:]), Rob(nums[:len(nums)-1]))
}
