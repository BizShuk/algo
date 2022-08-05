package backtrack

// [Pattern]: [Walking Map BT, DFS]
func WalkMap(board [][]byte, word string) bool {
	found := false

	rsize := len(board)
	csize := len(board[0])

	marked := make([][]bool, rsize)
	for i := 0; i < rsize; i++ {
		marked[i] = make([]bool, csize)
	}

	var validPosition func(i, j int) bool = func(i, j int) bool {
		if i < 0 || j < 0 || i >= rsize || j >= csize {
			return false
		}
		return true
	}
	var xshift []int = []int{1, -1, 0, 0}
	var yshift []int = []int{0, 0, 1, -1}

	var dfs func(x, y int, wi int) bool /* [Hint]: Change return type depended on questions */
	dfs = func(x, y int, wi int) bool {
		if !validPosition(x, y) || marked[x][y] {
			return false
		}
		// [Hint]: Put some conditions here

		marked[x][y] = true
		defer func() {
			marked[x][y] = false
		}()

		for i := 0; i < 4; i++ {
			dfs(x+xshift[i], y+yshift[i], wi+1)
			// [Hint]: Put some conditions here
		}

		return false
	}
	// for starting point[s]
	for i := 0; i < rsize; i++ {
		for j := 0; j < csize; j++ {
			if dfs(i, j, 0) {
				found = true
				return true
			}
		}
	}

	return found
}
