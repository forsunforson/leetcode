/*
 * @lc app=leetcode.cn id=227 lang=golang
 *
 * [227] 基本计算器 II
 */
package main

// @lc code=start
func calculate(s string) int {
	stack := make([]int, 0)
	num := 0
	preSig := '+'
	for idx, n := range s {
		if isDigit(n) {
			num = num*10 + int(n-'0')
		}
		// 遇到计算符号或者已经遍历完数组，都得进行计算
		if !isDigit(n) && n != ' ' || idx == len(s)-1 {
			switch preSig {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, top*num)
			case '/':
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, top/num)
			default:
			}
			num = 0
			preSig = n
		}

	}
	result := 0
	for _, n := range stack {
		result += n
	}
	return result
}

func isDigit(n rune) bool {
	return n >= '0' && n <= '9'
}

// @lc code=end
