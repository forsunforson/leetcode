/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长递增子序列
 */
package main

// @lc code=start
func lengthOfLIS(nums []int) int {
	maxLength := 1
	dp := make([]int, len(nums))
	dp[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		max := 1
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[i] {
				if max < dp[j]+1 {
					max = dp[j] + 1
				}
			}
		}
		dp[i] = max
		if maxLength < dp[i] {
			maxLength = max
		}
	}
	return maxLength
}

// @lc code=end
