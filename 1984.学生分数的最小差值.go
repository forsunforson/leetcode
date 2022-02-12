/*
 * @lc app=leetcode.cn id=1984 lang=golang
 *
 * [1984] 学生分数的最小差值
 */
package main

import "sort"

// @lc code=start
func minimumDifference(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	// 排序，然后滑动窗口取值
	sort.Ints(nums)
	minDiff := 100000
	l, r := 0, k-1
	for {
		if r >= len(nums) {
			break
		}
		minDiff = min(minDiff, abs(nums[l], nums[r]))
		l++
		r++
	}
	return minDiff
}

func abs(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end
