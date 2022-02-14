/*
 * @lc app=leetcode.cn id=650 lang=golang
 *
 * [650] 只有两个键的键盘
 */
package main

// @lc code=start
func minSteps(n int) int {
	// 不能输入A，只能复制粘贴
	if n == 1 {
		return 0
	}
	dp := make([]int, n+1)
	dp[1] = 0
	dp[2] = 2
	for i := 3; i < n+1; i++ {
		dp[i] = i
	}
	for i := 2; i < n+1; i++ {
		for j := 2; i*j < n+1; j++ {
			dp[i*j] = min(dp[i*j], dp[i]+j)
		}
	}
	return dp[n]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end
