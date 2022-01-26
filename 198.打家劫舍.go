/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */
package main

import "fmt"

// @lc code=start
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	// 动态规划，n的数字由n-1和n-2中大的决定
	var result []int
	result = append(result, 0, nums[0])
	for i := 1; i < len(nums); i++ {
		if result[i] > result[i-1]+nums[i] {
			result = append(result, result[i])
		} else {
			result = append(result, result[i-1]+nums[i])
		}

	}
	fmt.Printf("%v", result)
	return result[len(nums)]
}

// @lc code=end
