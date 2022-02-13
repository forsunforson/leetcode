/*
 * @lc app=leetcode.cn id=167 lang=golang
 *
 * [167] 两数之和 II - 输入有序数组
 */
package main

// @lc code=start
func twoSum(numbers []int, target int) []int {
	l, r := 1, len(numbers)
	for {
		sum := numbers[l-1] + numbers[r-1]
		if sum == target {
			break
		} else if sum < target {
			l++
		} else if sum > target {
			r--
		}
	}
	return []int{l, r}
}

// @lc code=end
