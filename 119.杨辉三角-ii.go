/*
 * @lc app=leetcode.cn id=119 lang=golang
 *
 * [119] 杨辉三角 II
 */
package main

// @lc code=start
func getRow(rowIndex int) []int {
	ans := make([]int, rowIndex+1)
	ans[0] = 1
	for i := 1; i <= rowIndex; i++ {
		for j := i; j > 0; j-- {
			ans[j] = ans[j] + ans[j-1]
		}
	}
	return ans
}

func sol2(rowIndex int) []int {
	ans := make([]int, rowIndex+1)
	ans[0] = 1
	for i := 1; i <= rowIndex; i++ {
		ans[i] = ans[i-1] * (rowIndex - i + 1) / i
	}
	return ans
}

// @lc code=end
