/*
 * @lc app=leetcode.cn id=415 lang=golang
 *
 * [415] 字符串相加
 */
package main

// @lc code=start
func addStrings(num1 string, num2 string) string {
	if len(num1) > len(num2) {
		return addStrings(num2, num1)
	}
	b1 := []byte(num1)
	b2 := []byte(num2)
	ans := make([]byte, len(num2)+1)
	carry := byte(0)
	for i := 1; i <= len(num1); i++ {
		sum := b1[len(num1)-i] - '0' + b2[len(num2)-i] - '0' + carry
		if sum > 9 {
			carry = 1
			sum -= 10
		} else {
			carry = 0
		}
		ans[len(ans)-i] = sum + '0'
	}
	for i := len(num1) + 1; i <= len(num2); i++ {
		sum := b2[len(num2)-i] - '0' + carry
		if sum > 9 {
			carry = 1
			sum -= 10
		} else {
			carry = 0
		}
		ans[len(ans)-i] = sum + '0'
	}
	if carry != 0 {
		ans[0] = 1 + '0'
		return string(ans)
	}
	return string(ans[1:])
}

// @lc code=end
