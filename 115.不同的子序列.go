/*
 * @lc app=leetcode.cn id=115 lang=golang
 *
 * [115] 不同的子序列
 */
package main

// @lc code=start
func numDistinct(s string, t string) int {
	if len(s) < len(t) {
		return 0
	}
	dp := make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, len(t)+1)
	}
	for i := 0; i <= len(s); i++ {
		dp[i][0] = 1
	}
	// 递推的过程是dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
	// 代表选了当前匹配的字符，或者不选，从之前匹配的中累加
	for i := 1; i < len(dp); i++ {
		for j := 1; j <= i && j < len(dp[i]); j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(s)][len(t)]
}

// @lc code=end
