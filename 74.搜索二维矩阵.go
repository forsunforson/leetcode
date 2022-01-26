/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */
package main

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) < 1 || len(matrix[0]) < 1 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	curRow := 0
	for ; curRow < m && matrix[curRow][n-1] <= target; curRow++ {
		if matrix[curRow][n-1] == target {
			return true
		}
	}
	if curRow == m {
		return false
	}
	begin, end := 0, n-1

	for {
		if begin > end {
			break
		}
		curCol := (begin + end) / 2
		if matrix[curRow][begin] == target || matrix[curRow][end] == target {
			return true
		}
		curNum := matrix[curRow][curCol]
		if curNum == target {
			return true
		}
		if curNum < target {
			begin = curCol + 1
		} else {
			end = curCol - 1
		}

	}
	return false
}

// @lc code=end
