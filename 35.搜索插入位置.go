/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */
package main

// @lc code=start
func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (r-l)/2 + l
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	if nums[l] >= target {
		return l
	} else {
		return l + 1
	}
}

// @lc code=end
