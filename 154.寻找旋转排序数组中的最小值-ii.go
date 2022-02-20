/*
 * @lc app=leetcode.cn id=154 lang=golang
 *
 * [154] 寻找旋转排序数组中的最小值 II
 */
package main

// @lc code=start
func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		// 对于l和r而言，nums[r]<=nums[l]
		mid := (r-l)/2 + l
		if nums[r] < nums[mid] {
			l = mid + 1
		} else if nums[r] > nums[mid] {
			r = mid
		} else {
			r--
		}
	}
	return nums[l]
}

// @lc code=end
