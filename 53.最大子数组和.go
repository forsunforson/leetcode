/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子数组和
 */
package main

import "math"

// @lc code=start
func maxSubArray(nums []int) int {
	maxSum := math.MinInt32
	sum := 0
	for _, num := range nums {
		sum += num
		maxSum = max(maxSum, sum)
		if sum < 0 {
			sum = 0
		}
	}
	return maxSum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end
