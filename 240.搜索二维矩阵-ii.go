/*
 * @lc app=leetcode.cn id=240 lang=golang
 *
 * [240] 搜索二维矩阵 II
 */
package main

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	x, y := 0, len(matrix[0])-1
	for x >= 0 && x < len(matrix) && y >= 0 && y < len(matrix[0]) {
		if matrix[x][y] == target {
			return true
		}
		if matrix[x][y] > target {
			y--
		} else {
			x++
		}
	}
	return false
}

// @lc code=end
