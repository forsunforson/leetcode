/*
 * @lc app=leetcode.cn id=1312 lang=golang
 *
 * [1312] 让字符串成为回文串的最少插入次数
 */
package main

// @lc code=start
func minInsertions(s string) int {
	// 找到最长回文子串，然后算差
	return len(s) - findLongestSubstr(s)
}

func findLongestSubstr(s string) int {
	// dp[n][n]
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	// 对于对角线上的值，是字符串本身，都为1
	// 答案在左上角，从右下角开始，横向扫描
	for i := n - 1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end
