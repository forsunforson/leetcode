/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 */
package main

// @lc code=start
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return max(robHelper(nums[:len(nums)-1]), robHelper(nums[1:]))
}

func robHelper(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		if i > 1 {
			dp[i] = max(nums[i]+dp[i-2], dp[i-1])
		} else {
			dp[i] = max(dp[i-1], nums[i])
		}
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
