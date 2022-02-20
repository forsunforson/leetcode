/*
 * @lc app=leetcode.cn id=97 lang=golang
 *
 * [97] 交错字符串
 */
package main

// @lc code=start
func isInterleave(s1 string, s2 string, s3 string) bool {
	// 当s1[i] == s2[j] == s3[k]时存在多种选择
	// dp[n][m] 表示s1[0:n]s2[0:m]时是否能拼出s3[0:n+m]
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	n, m := len(s1), len(s2)
	dp := make([]bool, m+1)
	dp[0] = true
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			k := i + j - 1
			if i > 0 {
				dp[j] = dp[j] && s1[i-1] == s3[k]
			}
			if j > 0 {
				dp[j] = dp[j] || dp[j-1] && s2[j-1] == s3[k]
			}
		}
	}
	return dp[m]
}

// @lc code=end
