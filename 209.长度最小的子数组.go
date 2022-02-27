/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 */
package main

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum < target {
		return 0
	}
	minLen := len(nums)
	l, r := 0, 0
	sum = 0
	for r < len(nums) {
		for r < len(nums) && sum < target {
			sum += nums[r]
			r++
		}

		for l < r && sum >= target {
			minLen = min(minLen, r-l)
			sum -= nums[l]
			l++
		}
	}
	return minLen
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end
