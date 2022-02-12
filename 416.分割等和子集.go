/*
 * @lc app=leetcode.cn id=416 lang=golang
 *
 * [416] 分割等和子集
 */
package main

import "sort"

// @lc code=start
func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	// 能将数组分成两个相等的子数组，和给数组加上正负号构成等式是一样的
	// 最直观的就是回溯法

	// 使用动态规划
	// dp[idx][result]
	sort.Ints(nums)
	n := len(nums)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, target+1)
	}
	dp[0][0] = true
	if nums[0] <= target {
		dp[0][nums[0]] = true
	}
	// 从上一行继承
	for i := 1; i < n; i++ {
		for j := 0; j <= target; j++ {
			if dp[i-1][j] {
				dp[i][j] = true
				if j+nums[i] <= target {
					dp[i][j+nums[i]] = true
				}
			}
		}
	}
	return dp[n-1][target]
}

// 回溯可以解，但是会超时
func findTarget(target int, nums []int, idx int) bool {
	if target == 0 {
		return true
	}
	if target < 0 {
		return false
	}
	if idx >= len(nums) {
		return false
	}
	return findTarget(target-nums[idx], nums, idx+1) ||
		findTarget(target, nums, idx+1)
}

// @lc code=end
