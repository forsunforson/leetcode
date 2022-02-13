/*
 * @lc app=leetcode.cn id=503 lang=golang
 *
 * [503] 下一个更大元素 II
 */
package main

// @lc code=start
func nextGreaterElements(nums []int) []int {
	// 拼接，循环
	helper := nums
	helper = append(helper, nums...)
	ans := []int{}
	stack := []int{}
	for i := len(helper) - 1; i >= 0; i-- {
		num := helper[i]
		for len(stack) > 0 && stack[len(stack)-1] <= num {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			ans = append(ans, -1)
		} else {
			ans = append(ans, stack[len(stack)-1])
		}
		stack = append(stack, num)
	}
	result := []int{}
	for i := 0; i < len(nums); i++ {
		result = append(result, ans[len(ans)-i-1])
	}
	return result
}

// @lc code=end
