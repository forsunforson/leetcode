/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 */
package main

// @lc code=start
func minimumTotal(triangle [][]int) int {
	// 倒序 dp
	dp := make([][]int, len(triangle))
	for i := range dp {
		dp[i] = make([]int, len(triangle[i]))
		if i == len(dp)-1 {
			dp[i] = triangle[i]
		}
	}
	for i := len(dp) - 2; i >= 0; i-- {
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j]
		}
	}
	return dp[0][0]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end
