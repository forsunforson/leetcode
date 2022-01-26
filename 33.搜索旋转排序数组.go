/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */
package main

// @lc code=start
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	// 先找到旋转点，再二分查找
	point := -1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			point = i
			break
		}
	}
	left, right := 0, len(nums)-1
	if point != -1 {
		if target < nums[0] {
			left = point
		} else {
			right = point - 1
		}
	}

	for {
		if left > right {
			break
		}
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

// @lc code=end
