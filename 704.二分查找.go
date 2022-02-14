/*
 * @lc app=leetcode.cn id=704 lang=golang
 *
 * [704] 二分查找
 */
package main

// @lc code=start
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (r-l)/2 + l
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		}
	}
	if l < len(nums) && nums[l] == target {
		return l
	}
	return -1
}

// @lc code=end
