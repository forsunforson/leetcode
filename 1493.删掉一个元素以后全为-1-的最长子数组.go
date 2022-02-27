/*
 * @lc app=leetcode.cn id=1493 lang=golang
 *
 * [1493] 删掉一个元素以后全为 1 的最长子数组
 */
package main

import "fmt"

// @lc code=start
func longestSubarray(nums []int) int {
	// 统计连续1的长度，并记录数组i前0的个数
	ones := []int{}
	zeros := []int{}
	cur := 0
	for idx := 0; idx < len(nums); {
		count := 0
		for idx < len(nums) && nums[idx] == cur {
			idx++
			count++
		}
		if cur == 0 {
			zeros = append(zeros, count)
		} else {
			ones = append(ones, count)
		}
		if idx < len(nums) {
			cur = nums[idx]
		} else {
			break
		}
	}
	fmt.Printf("1: %v 0: %v", ones, zeros)
	if len(ones) == 0 {
		return 0
	}
	if len(ones) == 1 && len(zeros) == 1 && zeros[0] == 0 {
		return ones[0] - 1
	}
	maxL := ones[0]
	for i := 1; i < len(ones); i++ {
		if zeros[i] == 1 {
			maxL = max(maxL, ones[i-1]+ones[i])
		}
		maxL = max(maxL, ones[i])
	}
	return maxL
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end
