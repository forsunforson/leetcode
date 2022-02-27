/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 */
package main

// @lc code=start
func generate(numRows int) [][]int {
	ans := [][]int{{1}}
	for i := 1; i < numRows; i++ {
		row := []int{1}
		for j := 0; j < len(ans[i-1]); j++ {
			add := 0
			if j+1 < len(ans[i-1]) {
				add = ans[i-1][j+1]
			}
			row = append(row, ans[i-1][j]+add)
		}
		ans = append(ans, row)
	}
	return ans
}

// @lc code=end
