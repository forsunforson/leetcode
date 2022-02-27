/*
 * @lc app=leetcode.cn id=343 lang=golang
 *
 * [343] 整数拆分
 */
package main

// @lc code=start
func integerBreak(n int) int {
	// 看起来有点像剪绳子
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3
	for i := 4; i <= n; i++ {
		dp[i] = max(dp[i-2]*2, dp[i-3]*3)
	}
	return dp[n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end
