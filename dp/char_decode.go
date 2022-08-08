package dp

// [Notice]: [Combination DP] Number Decode for alphabet
// edge case:
// 27 28 29(take one/two for 2, take one for 7 8 9)
// 30 40 50 60 70 80 90(take two)
// 3,4,5,6,7,8,9(take two with before, or take one)
func NumDecodings(s string) int {
	countArea := []int{}

	i := 0
	c1, c2 := 0, 1

	for i < len(s) {
		if s[i] == '0' { // leading zero
			return 0
		}

		if i+1 < len(s) && s[i+1] == '0' && string(s[i:i+2]) > "20" { // 30 40 50 60 70 80 90 can't decode
			return 0
		}

		if i+1 < len(s) && (string(s[i:i+2]) == "10" ||
			string(s[i:i+2]) == "20") { // 10 20

			i += 2
			countArea = append(countArea, c2, 1)
			c1, c2 = 0, 1
			continue
		}

		if i+1 < len(s) && (string(s[i:i+2]) == "27" ||
			string(s[i:i+2]) == "28" ||
			string(s[i:i+2]) == "29") { // 27 28 29

			i += 2
			countArea = append(countArea, c2+c1)
			c1, c2 = 0, 1
			continue
		}

		if s[i] >= '3' && s[i] <= '9' { // 3 4 5 6 7 8 9 ,  27 28 29
			i += 1
			countArea = append(countArea, c2+c1)
			c1, c2 = 0, 1
			continue
		}

		c1, c2 = c2, c1+c2
		i += 1
	}

	countArea = append(countArea, c2)

	count := 1

	for _, num := range countArea {
		count *= num
	}
	return count
}
