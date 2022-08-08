package array

// [Notice]: column index of grid is col/3, row index of grid is (row/3)*3. https://leetcode.com/problems/valid-sudoku/

func isValidSudoku(board [][]byte) bool {

	rows := make([][9]bool, 9)
	cols := make([][9]bool, 9)
	grids := make([][9]bool, 9)

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			num := board[row][col]

			if num == '.' {
				continue
			}
			num = num - 49
			k := (row/3)*3 + (col / 3)
			// fmt.Println(num+1, k , row, col)

			if rows[row][num] || cols[col][num] || grids[k][num] {
				return false
			}

			rows[row][num] = true
			cols[col][num] = true
			grids[k][num] = true
		}
	}
	return true
}
