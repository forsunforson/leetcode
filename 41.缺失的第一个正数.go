/*
 * @lc app=leetcode.cn id=41 lang=golang
 *
 * [41] 缺失的第一个正数
 */
package main

// @lc code=start
func firstMissingPositive(nums []int) int {
	// 未出现的正数范围[1, len(nums)+1]
	n := len(nums)
	for i := range nums {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}
	for _, num := range nums {
		num = abs(num)
		if num <= n {
			nums[num-1] = -abs(nums[num-1])
		}
	}
	for i, num := range nums {
		if num > 0 {
			return i + 1
		}
	}
	return n + 1
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

// @lc code=end
