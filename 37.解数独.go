/*
 * @lc app=leetcode.cn id=37 lang=golang
 *
 * [37] 解数独
 */
package main

// @lc code=start
func solveSudoku(board [][]byte) {
	// 用回溯法遍历
	// 和八皇后类似，需要三个map判断是否合法
	var row, col [9][9]bool
	var block [3][3][9]bool
	// 不同点，在于需要判断还有多少个空格和位置，来决定退出
	var spaces [][2]int
	// 初始化map
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := board[i][j] - '1'
				row[i][num] = true
				col[j][num] = true
				block[i/3][j/3][num] = true
			} else {
				spaces = append(spaces, [2]int{i, j})
			}
		}
	}
	var dfs func(int) bool
	dfs = func(pos int) bool {
		if pos == len(spaces) {
			return true
		}
		i, j := spaces[pos][0], spaces[pos][1]
		for n := 0; n < 9; n++ {
			if !row[i][n] && !col[j][n] && !block[i/3][j/3][n] {
				row[i][n] = true
				col[j][n] = true
				block[i/3][j/3][n] = true
				board[i][j] = byte(n) + '1'
				if dfs(pos + 1) {
					return true
				}
				board[i][j] = '.'
				row[i][n] = false
				col[j][n] = false
				block[i/3][j/3][n] = false
			}
		}
		return false
	}
	dfs(0)
}

// @lc code=end
