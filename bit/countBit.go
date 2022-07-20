package bit

func CountBit(num int) []int {
	dp := make([]int, num+1)
	dp[0] = 0

	offset := 1

	for i := 1; i < num+1; i++ {
		if offset<<1 == i {
			offset = i
		}
		dp[i] = 1 + dp[i-offset]

	}

	return dp
}
