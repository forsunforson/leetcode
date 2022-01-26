/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */
package main

// @lc code=start
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			visited := make([][]bool, m)
			for k := 0; k < m; k++ {
				visited[k] = make([]bool, n)
			}
			if existHelper(board, word, &visited, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func existHelper(board [][]byte, word string, visited *[][]bool, x, y, idx int) bool {
	if idx == len(word) {
		return true
	}
	if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) {
		return false
	}
	if (*visited)[x][y] {
		return false
	}

	if board[x][y] != word[idx] {

		return false
	}
	(*visited)[x][y] = true
	result := existHelper(board, word, visited, x+1, y, idx+1) ||
		existHelper(board, word, visited, x-1, y, idx+1) ||
		existHelper(board, word, visited, x, y+1, idx+1) ||
		existHelper(board, word, visited, x, y-1, idx+1)
	if result {
		return true
	}
	(*visited)[x][y] = false
	return false
}

// @lc code=end
