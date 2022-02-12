/*
 * @lc app=leetcode.cn id=130 lang=golang
 *
 * [130] 被围绕的区域
 */

// @lc code=start
func solve(board [][]byte) {
	// 只有和边界相连的O才不会被X包围
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || i == m-1 || j == 0 || j == n-1 {
				tagBoarder(&board, i, j)
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
}

func tagBoarder(board *[][]byte, x, y int) {
	if x < 0 || x >= len(*board) || y < 0 || y >= len((*board)[0]) {
		return
	}
	if (*board)[x][y] != 'O' {
		return
	}
	(*board)[x][y] = '#'
	tagBoarder(board, x+1, y)
	tagBoarder(board, x-1, y)
	tagBoarder(board, x, y+1)
	tagBoarder(board, x, y-1)
}

// @lc code=end

