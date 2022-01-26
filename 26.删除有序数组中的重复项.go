/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */
package main

// @lc code=start
func removeDuplicates(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	writeIdx := -1
	for _, num := range nums {
		if writeIdx == -1 {
			writeIdx++
			continue
		}
		if num == nums[writeIdx] {
			continue
		} else {
			writeIdx++
			nums[writeIdx] = num
		}
	}
	return writeIdx + 1
}

// @lc code=end
