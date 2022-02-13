/*
 * @lc app=leetcode.cn id=224 lang=golang
 *
 * [224] 基本计算器
 */
package main

// @lc code=start
func calculate(s string) int {
	// 使用两个栈来存储
	// 一个栈存储数字；一个栈存储括号
	ans := 0
	ops := []int{1}
	sign := 1
	for i := 0; i < len(s); {
		switch s[i] {
		case ' ':
			i++
		case '+':
			sign = ops[len(ops)-1]
			i++
		case '-':
			sign = -ops[len(ops)-1]
			i++
		case '(':
			ops = append(ops, sign)
			i++
		case ')':
			ops = ops[:len(ops)-1]
			i++
		default:
			num := 0
			for ; i < len(s) && '0' <= s[i] && s[i] <= '9'; i++ {
				num = 10*num + int(s[i]-'0')
			}
			ans += sign * num
		}
	}
	return ans
}

// @lc code=end
