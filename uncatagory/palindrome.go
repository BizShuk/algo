package uncatagory

// Notice:
// byte(uint8) is from s[i]
// rune(int32) is from for range value

// Two pointers solutions
// Fast and slow solution: reverse the 2nd half and use faster pointer/head pointer to compare

func IsPalindrome(s string) bool {
	size := len(s)
	f, l := 0, size

	for f < l {
		if s[f] != s[l] {
			return false
		}
		f++
		l--
	}
	return true
}
