/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */
package main

// @lc code=start
// 正序，算出0的个数
func moveZeroes(nums []int) {
	countZero := 0
	idx := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[idx] = nums[i]
			idx++
		} else {
			countZero++
		}
	}
	for i := len(nums) - countZero; i < len(nums); i++ {
		nums[i] = 0
	}
}

// @lc code=end
