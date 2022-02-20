/*
 * @lc app=leetcode.cn id=334 lang=golang
 *
 * [334] 递增的三元子序列
 */
package main

import "math"

// @lc code=start
func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	l, r := nums[0], math.MaxInt32
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num > r {
			return true
		} else if num > l {
			r = num
		} else {
			l = num
		}
	}
	return false
}

// @lc code=end
