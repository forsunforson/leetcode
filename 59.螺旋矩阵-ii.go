/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 */
package main

// @lc code=start
func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	l, r, t, b := 0, len(matrix[0]), 0, len(matrix)
	var write func(int, int, int, int)
	count := 1
	write = func(l, r, t, b int) {
		for i := l; i < r; i++ {
			matrix[t][i] = count
			count++
		}
		for i := t + 1; i < b-1; i++ {
			matrix[i][r-1] = count
			count++
		}
		if b-t > 1 {
			for i := r - 1; i >= l; i-- {
				matrix[b-1][i] = count
				count++
			}
		}
		if r-l > 1 {
			for i := b - 2; i >= t+1; i-- {
				matrix[i][l] = count
				count++
			}
		}
	}
	for l < r && t < b {
		write(l, r, t, b)
		l++
		r--
		t++
		b--
	}
	return matrix
}

// @lc code=end
