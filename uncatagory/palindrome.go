package uncatagory

/*
	treat input string a pure acsii characters
*/

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
