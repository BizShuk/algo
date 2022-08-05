package dp

import "github.com/bizshuk/algo/util"

// Longest Common Subsequence

// [Pattern]: [Longest Common Subsequence DP] find longest common subsequence in two strings.
// abcde, acz => ac => length = 2
func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}

	for i := len(text1) - 1; i >= 0; i-- {
		for j := len(text2) - 1; j >= 0; j-- {
			if text1[i] == text2[j] {
				dp[i][j] = 1 + dp[i+1][j+1]
			} else {
				dp[i][j] = util.Max(dp[i+1][j], dp[i][j+1])
			}
		}
	}
	return dp[0][0]
}
