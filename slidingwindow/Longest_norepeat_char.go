package slidingwindow

// [Notice]: Sliding Window with left/right pointers
func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	set := make(map[byte]int)
	l, r := 0, 0
	for r = 0; r < len(s); r++ {
		for set[s[r]] > 0 {
			delete(set, s[l])
			l += 1
		}

		set[s[r]] += 1
		maxLength = Max(maxLength, r-l+1)
	}
	return maxLength
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
