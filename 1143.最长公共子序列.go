/*
 * @lc app=leetcode.cn id=1143 lang=golang
 *
 * [1143] 最长公共子序列
 */
package main

// @lc code=start
func longestCommonSubsequence(text1 string, text2 string) int {
	// dp[n][m]
	n, m := len(text1), len(text2)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[n][m]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end
