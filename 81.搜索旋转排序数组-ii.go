/*
 * @lc app=leetcode.cn id=81 lang=golang
 *
 * [81] 搜索旋转排序数组 II
 */
package main

// @lc code=start
func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (r-l)/2 + l
		if nums[mid] == target {
			return true
		}
		if nums[l] == nums[mid] && nums[mid] == nums[r] {
			l++
			r--
			continue
		}
		if nums[mid] >= nums[l] {
			// 前半按顺序
			if nums[mid] > target && target >= nums[l] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			// 后半有序
			if nums[mid] < target && nums[r] >= target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return false
}

// @lc code=end
