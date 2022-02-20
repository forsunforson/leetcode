/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 */
package main

// @lc code=start
func findMin(nums []int) int {
	// 和之前的题目一样，一定有一半是有序的
	l, r := 0, len(nums)-1
	for l < r {
		mid := (r-l)/2 + l
		if nums[l] <= nums[mid] && nums[mid] <= nums[r] {
			// 当前有序
			return nums[l]
		}
		if nums[l] <= nums[mid] {
			// 左半有序
			l = mid + 1
		} else {
			r = mid
		}
	}
	return nums[l]
}

// @lc code=end
