/*
 * @lc app=leetcode.cn id=162 lang=golang
 *
 * [162] 寻找峰值
 */
package main

import "math"

// @lc code=start
func findPeakElement(nums []int) int {
	// 也可以根据爬坡法做，如果mid-1 < mid < mid+1，那么选右区间;否则左区间
	l, r := 0, len(nums)-1
	var dfs func(int, int) int
	dfs = func(l, r int) int {
		mid := (l + r) / 2
		lNum, rNum := math.MinInt64, math.MinInt64
		if mid > 0 {
			lNum = nums[mid-1]
		}
		if mid < len(nums)-1 {
			rNum = nums[mid+1]
		}
		if nums[mid] > lNum && nums[mid] > rNum {
			return mid
		}
		if mid > l {
			left := dfs(l, mid-1)
			if left != -1 {
				return left
			}
		}
		if mid < r {
			right := dfs(mid+1, r)
			return right
		}
		return -1
	}
	return dfs(l, r)
}

// @lc code=end
