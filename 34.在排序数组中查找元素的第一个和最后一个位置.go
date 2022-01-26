/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */
package main

// @lc code=start
func searchRange(nums []int, target int) []int {
	// 二分法做两次
	left, right := 0, len(nums)-1
	l, r := -1, -1
	for {
		if left > right {
			break
		}
		mid := (left + right) / 2
		if nums[mid] == target {
			if mid == left || nums[mid-1] != nums[mid] {
				l = mid
				break
			} else {
				right = mid - 1
			}
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	left, right = 0, len(nums)-1
	for {
		if left > right {
			break
		}
		mid := (left + right) / 2
		if nums[mid] == target {
			if mid == right || nums[mid+1] != nums[mid] {
				r = mid
				break
			} else {
				left = mid + 1
			}
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return []int{l, r}
}

// @lc code=end
