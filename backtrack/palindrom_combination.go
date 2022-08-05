package backtrack

// [Variant]: [Palindrom DP] list all palindrom combination of a string
// aab => [[a,a,b],[aa,b]]
func Partition(s string) [][]string {
	palindromMatrix := make([][]bool, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		palindromMatrix[i] = make([]bool, len(s))

		for j := 0; i+j < len(s); j++ {
			if j == 0 {
				palindromMatrix[i][j] = true
				continue
			}

			if j == 1 && s[i] == s[i+j] {
				palindromMatrix[i][j] = true
				continue
			}

			if j == 1 {
				continue
			}

			// fmt.Println(i,j, palindromMatrix)
			if palindromMatrix[i+1][j-2] && s[i] == s[i+j] {
				palindromMatrix[i][j] = true
				continue
			}
		}
	}

	// fmt.Println(palindromMatrix)
	result := [][]string{}
	stack := []string{}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i == len(s) && j == 0 {
			result = append(result, stack)
			return
		}

		if i+j >= len(s) {
			return
		}

		dfs(i, j+1)

		if s[i] == s[i+j] && palindromMatrix[i][j] {
			stack = append(stack, string(s[i:i+j+1]))
			dfs(i+j+1, 0)
			stack = stack[: len(stack)-1 : len(stack)-1]
		}
	}

	dfs(0, 0)
	return result
}
