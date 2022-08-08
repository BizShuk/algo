package stack

// => How many path could a tree have?
// [Notice]: Generate combination "(())" from top-down [unknown tree path]

var size int

func generateParenthesis2(n int) []string {
	size = n
	return dfs2([]byte{}, 0, 0)
}

func dfs2(b []byte, i, j int) []string {
	if i == size && j == size {
		return []string{string(b)}
	}
	var s1, s2 []string

	if i < size {
		b = append(b, '(')
		s1 = dfs2(b, i+1, j)
		b = b[:len(b)-1]
	}

	if i > j && size > j {
		b = append(b, ')')
		s2 = dfs2(b, i, j+1)
		b = b[:len(b)-1]
	}
	return append(s1, s2...)
}

func generateParenthesis(n int) []string {
	stack := []string{}

	var dfs func(s string, i, j int)

	dfs = func(s string, i, j int) {
		if i == 0 && j == 0 {
			stack = append(stack, s)
		}

		if i > 0 {
			dfs(s+"(", i-1, j)
		}

		if j > i {
			dfs(s+")", i, j-1)
		}
	}

	dfs("", n, n)

	return stack
}
