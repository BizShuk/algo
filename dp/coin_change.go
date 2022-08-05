package dp

import "github.com/bizshuk/algo/util"

// [Pattern]: [Combination DP] How many ways of Coin Change
// [DP solution]: bottom-up with memorization
func CoinChange(coins []int, amount int) int {
	// 0 at index 0 is valid
	// 0 at index non-zero is invalid(no change way)
	dp := make([]int, amount+1)

	if amount == 0 {
		return 0
	}
	// O(N*M), N: amount, M: num of coins
	for i := 1; i <= amount; i++ {
		min := 100005
		for _, c := range coins {
			if i < c { // can't change
				continue
			}

			if i == c { // avoid use amount 0 to calculate ambiguous case
				dp[i] = 1
				break
			}

			if dp[i-c] == 0 { // for those at index non-zero, invalid changes
				continue
			}
			min = util.Min(min, dp[i-c])
			dp[i] = min + 1
		}
	}

	if dp[amount] == 0 {
		return -1
	}

	return dp[amount]
}

// [Variant]: [Combination BT]
func CoinChange_recursive(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	min := 100005
	for _, c := range coins {
		if amount-c == 0 {
			return 1
		}
		if amount-c < 0 {
			continue
		}

		change := CoinChange_recursive(coins, amount-c)

		if change == -1 {
			continue
		}

		min = util.Min(min, change)
	}

	if min == 100005 {
		return -1
	}
	return min + 1
}

// [Variant]: [Combination DP] Could Work Break by dictionary
func WordBreak_BU(s string, wordDict []string) bool {
	if s == "" {
		return true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := range s {
		dpi := i + 1
		for _, word := range wordDict {
			wSize := len(word)
			if dpi-wSize < 0 {
				continue
			}

			if dp[dpi-wSize] && string(s[dpi-wSize:dpi]) == word {
				dp[dpi] = true
			}

		}

	}
	return dp[len(s)]
}

// Top-Down, (duplication calculation)
func WordBreak_TD(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}

	for _, word := range wordDict {
		wSize := len(word)
		if wSize > len(s) {
			continue
		}

		if word != string(s[:wSize]) {
			continue
		}

		valid := WordBreak_TD(string(s[wSize:]), wordDict)
		if valid {
			return true
		}
	}
	return false
}
