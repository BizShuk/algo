package dp

// Center point solution: O(n^2)
func CountPalindromSubstrings(s string) int {
	count := 0
	for i := range s {
		count += CountPalindrom(s, i, i)
		count += CountPalindrom(s, i, i+1)
	}

	return count
}

func CountPalindrom(s string, l int, r int) int {
	count := 0
	for 0 <= l && r < len(s) {
		if s[l] != s[r] {
			break
		}
		count += 1
		l, r = l-1, r+1
	}
	return count
}

// Notice:
// [Pattern]: [Palindrom DP]
// Check each substring by
// 	  1. First and last char of substring are equal. s[i] == s[i+l]
// 	  2. Narrower substring is palindrom.
// Eaxmple: bacab , i = 1, l = 2, dp[1][2]
// => dp[1][2] = true if
//   - s[i] (1st a) ==  s[i+l] (2nd a)
//   - dp[i+1][l-2] = dp[2][0] == true
func CountPalindromSubstrings_DP(s string) int {
	count := 0

	palinDP := make([][]bool, len(s))

	for i := range palinDP {
		palinDP[i] = make([]bool, len(s))
	}

	for i := 0; i < len(s); i++ { // Single char always palindrom
		palinDP[i][0] = true
		count++
	}

	for l := 1; l < len(s); l++ { // i:index, l:length ex: "abcde", i=1, l=2 => "bcd"
		for i := 0; i+l < len(s); i++ {
			if l == 1 && s[i] == s[i+1] {
				palinDP[i][l] = true
				count++
			} else if s[i] == s[i+l] && palinDP[i+1][l-2] {
				palinDP[i][l] = true
				count++
			}
		}
	}

	return count
}

// [Variant]: [Palindrom DP] Finding longest palindrom substring
func LongestPalindromeInSubstring(s string) string {
	longestSub := string(s[0])
	maxL := 1
	palinDP := make([][]bool, len(s))

	for i := range palinDP {
		palinDP[i] = make([]bool, len(s))
	}

	for i := 0; i < len(s); i++ { // Single char always palindrom
		palinDP[i][0] = true
	}

	for l := 1; l < len(s); l++ { // i:index, l:length ex: "abcde", i=1, l=2 => "bcd"
		for i := 0; i+l < len(s); i++ {
			if l == 1 && s[i] == s[i+1] {
				palinDP[i][l] = true
			} else if s[i] == s[i+l] && palinDP[i+1][l-2] {
				palinDP[i][l] = true
			}

			if palinDP[i][l] && maxL < l+1 {
				maxL = l + 1
				longestSub = s[i : i+l+1]
			}
		}
	}

	return longestSub
}
