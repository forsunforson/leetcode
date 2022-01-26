/*
 * @lc app=leetcode.cn id=287 lang=golang
 *
 * [287] 寻找重复数
 */
package main

// @lc code=start
func findDuplicate(nums []int) int {
	// 快慢指针
	fast, slow := 0, 0
	for {
		fast = nums[nums[fast]]
		slow = nums[slow]
		if fast == slow {
			break
		}
	}
	slow = 0
	for {
		if fast == slow {
			break
		}
		fast = nums[fast]
		slow = nums[slow]
	}
	return fast
}

// @lc code=end
