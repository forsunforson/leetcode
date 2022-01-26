/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */
package main

import "sort"

// @lc code=start

func nextPermutation(nums []int) {
	// 从右到左找
	// 找到第一个非升序的
	left := len(nums) - 1
	max := nums[left]
	for ; left >= 0 && nums[left] >= max; left-- {
		max = nums[left]
	}
	if left == -1 {
		sort.Ints(nums)
		return
	}
	// 找到left右边第一个比它大的，交换
	right := len(nums) - 1
	for ; nums[right] <= nums[left]; right-- {

	}
	nums[left], nums[right] = nums[right], nums[left]
	sort.Ints(nums[left+1:])
}

// @lc code=end
