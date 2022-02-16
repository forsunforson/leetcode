/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子数组
 */
package main

// @lc code=start
func maxProduct(nums []int) int {
	maxRes, minRes, ans := nums[0], nums[0], nums[0]
	for i, num := range nums {
		if i == 0 {
			continue
		}
		mx, mn := maxRes, minRes
		maxRes = max(mx*num, max(num, mn*num))
		minRes = min(mn*num, min(num, mx*num))
		ans = max(ans, maxRes)
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end
