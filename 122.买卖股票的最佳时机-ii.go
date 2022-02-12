/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start
func maxProfit(prices []int) int {
	// dp[i]表示当前天数获得的最大收益
	n := len(prices)
	dp := make([]int, n)
	// i-1 存在两种可能，持有股票或者持有现金
	for i := 1; i < n; i++ {
		dp[i] = dp[i-1] + max(0, prices[i]-prices[i-1])
	}
	return dp[n-1]
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// @lc code=end

