/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 */
package main

// @lc code=start
func longestValidParentheses(s string) int {
	return longestValidParenthesesPtr(s)
}

func longestValidParenthesesPtr(s string) int {
	// 双指针
	left, right := 0, 0
	maxLength := 0
	// 从左往右
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxLength = max(maxLength, 2*right)
		} else if right > left {
			// right 大于 left说明右括号太多了，得重新匹配
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	// 从右往左
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxLength = max(maxLength, 2*left)
		} else if left > right {
			left, right = 0, 0
		}
	}
	return maxLength
}

func longestValidParenthesesStack(s string) int {
	stack := make([]int, 0)
	stack = append(stack, -1)
	maxLenth := 0
	// 放入没有匹配的右括号时，更新栈
	// 确保栈底放着一个没有匹配的右括号
	for i, r := range s {
		if r == '(' {
			stack = append(stack, i)
		} else {
			// 直接把头给pop出来
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				maxLenth = max(maxLenth, i-stack[len(stack)-1])
			}
		}
	}
	return maxLenth
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// @lc code=end
