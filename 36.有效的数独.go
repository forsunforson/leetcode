/*
 * @lc app=leetcode.cn id=36 lang=golang
 *
 * [36] 有效的数独
 */
package main

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	var row, col [9][9]bool
	var block [3][3][9]bool
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			n := board[i][j] - '1'
			if !row[i][n] && !col[j][n] && !block[i/3][j/3][n] {
				row[i][n] = true
				col[j][n] = true
				block[i/3][j/3][n] = true
			} else {
				return false
			}
		}
	}
	return true
}

// @lc code=end
