/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */
package main

// @lc code=start
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	ptr := len(nums) - 1
	for ptr >= 0 && nums[ptr] == val {
		ptr--
	}
	l := 0
	for l < len(nums) && ptr >= 0 && l <= ptr {
		if nums[l] == val {
			nums[l], nums[ptr] = nums[ptr], nums[l]
			for nums[ptr] == val {
				ptr--
			}
		}
		l++
	}
	return l
}

// @lc code=end
