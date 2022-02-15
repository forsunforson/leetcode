/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */
package main

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	ans := []int{}
	l, r, t, b := 0, len(matrix[0]), 0, len(matrix)
	var printOut func(int, int, int, int) []int
	printOut = func(l, r, t, b int) []int {
		ret := []int{}
		for i := l; i < r; i++ {
			ret = append(ret, matrix[t][i])
		}
		for i := t + 1; i < b-1; i++ {
			ret = append(ret, matrix[i][r-1])
		}
		if b-t > 1 {
			for i := r - 1; i >= l; i-- {
				ret = append(ret, matrix[b-1][i])
			}
		}
		if r-l > 1 {
			for i := b - 2; i >= t+1; i-- {
				ret = append(ret, matrix[i][l])
			}
		}

		return ret
	}
	for l < r && t < b {
		ans = append(ans, printOut(l, r, t, b)...)
		l++
		r--
		t++
		b--
	}
	return ans
}

// @lc code=end
