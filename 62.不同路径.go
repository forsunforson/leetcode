/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */
package main

// @lc code=start
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[m-1][n-1] = 1
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			cur := dp[i][j]
			if i+1 < m {
				cur += dp[i+1][j]
			}
			if j+1 < n {
				cur += dp[i][j+1]
			}
			dp[i][j] = cur
		}
	}
	return dp[0][0]
}

// @lc code=end
