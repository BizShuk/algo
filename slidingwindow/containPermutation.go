package slidingwindow

// [Notice]: When moving left pointer, usually it should less than right pointer
// [Notice]: Where to ovve left pointer?
// [Notice]: Right pointer loop constraint. usually less than len(input)
// [Notice]: flow: since s2 is fixed, Using length of s1 as window move in s2

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	s1map := make([]byte, 26)
	s2map := make([]byte, 26)

	for i := 0; i < len(s1); i++ {
		s1map[s1[i]-97] += 1
		s2map[s2[i]-97] += 1
	}

	s1Key := string(s1map)

	l, r := 0, len(s1) // window length

	for ; r < len(s2); l, r = l+1, r+1 {
		if string(s2map) == s1Key {
			return true
		}

		// l and r move in the same time
		s2map[s2[r]-97] += 1
		s2map[s2[l]-97] -= 1
	}

	return string(s2map) == s1Key
}
