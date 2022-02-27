/*
 * @lc app=leetcode.cn id=80 lang=golang
 *
 * [80] 删除有序数组中的重复项 II
 */
package main

// @lc code=start
func removeDuplicates(nums []int) int {
	idx := 1
	cur := nums[0]
	curCount := 1
	for i := 1; i < len(nums); {
		if nums[i] == cur {
			curCount++
			if curCount > 2 {
				for i < len(nums) && nums[i] == cur {
					i++
				}
				if i >= len(nums) {
					break
				}
				cur = nums[i]
				curCount = 1
			}
		} else {
			cur = nums[i]
			curCount = 1
		}
		nums[idx] = nums[i]
		idx++
		i++
	}
	return idx
}

// @lc code=end
