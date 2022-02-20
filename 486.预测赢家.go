/*
 * @lc app=leetcode.cn id=486 lang=golang
 *
 * [486] 预测赢家
 */
package main

// @lc code=start
func PredictTheWinner(nums []int) bool {
	// 当前步需要最小化对手的得分
	// 对于玩家12来说，站在每一步的最优解都是相同的
	// dp[i][j]表示数组从[i, j]的最优解，是max(nums[i]-dp[i+1][j], nums[j]-dp[i][j-1])
	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, len(nums))
		// 数组长度只有1，肯定先手的赢
		dp[i][i] = nums[i]
	}

	for i := len(nums) - 1; i >= 0; i-- {
		for j := i + 1; j < len(nums); j++ {
			dp[i][j] = max(nums[i]-dp[i+1][j], nums[j]-dp[i][j-1])
		}
	}
	return dp[0][len(nums)-1] >= 0
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end
