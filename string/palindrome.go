package string

// [Variant]: [TwoPointers] for palindrome
// [Alternative] Fast and slow solution: reverse the 2nd half and use faster pointer/head pointer to compare
func IsPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r { // equal is not necessary since it should be identical
		if s[l] != s[r] {
			return false
		}
		l = l + 1
		r = r - 1
	}
	return true
}
