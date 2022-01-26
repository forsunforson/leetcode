/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */
package main

// @lc code=start
func canJump(nums []int) bool {
	// 倒推
	lastIdx := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i]+i >= lastIdx {
			lastIdx = i
		}
	}
	return lastIdx == 0
}

// @lc code=end
