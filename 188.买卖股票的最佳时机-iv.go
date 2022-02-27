/*
 * @lc app=leetcode.cn id=188 lang=golang
 *
 * [188] 买卖股票的最佳时机 IV
 */
package main

// @lc code=start
func maxProfit(k int, prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, len(prices))
	}
	for i := 1; i <= k; i++ {
		for j := 0; j < len(prices); j++ {
			dp[i][j] = dp[i-1][j]
			if j > 0 {
				dp[i][j] = max(dp[i][j], dp[i][j-1])
			}
			for k := 0; k < j; k++ {
				dp[i][j] = max(dp[i][j], dp[i-1][k]+prices[j]-prices[k])
			}
		}
	}
	return dp[k][len(prices)-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end
