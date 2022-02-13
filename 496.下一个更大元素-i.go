/*
 * @lc app=leetcode.cn id=496 lang=golang
 *
 * [496] 下一个更大元素 I
 */
package main

// @lc code=start
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	// 更大的元素指按当前排序下，x在nums2中的下一个大元素
	// 使用单调栈
	ans := []int{}
	stack := []int{}
	mapping := map[int]int{}
	for i := len(nums2) - 1; i >= 0; i-- {
		num := nums2[i]
		for len(stack) > 0 && stack[len(stack)-1] < num {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			mapping[num] = -1
		} else {
			mapping[num] = stack[len(stack)-1]
		}
		stack = append(stack, num)
	}
	for _, num := range nums1 {
		ans = append(ans, mapping[num])
	}
	return ans
}

// @lc code=end
