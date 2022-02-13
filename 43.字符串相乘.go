/*
 * @lc app=leetcode.cn id=43 lang=golang
 *
 * [43] 字符串相乘
 */
package main

import "fmt"

// @lc code=start
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	ans := "0"
	for i := len(num2) - 1; i >= 0; i-- {
		cur := ""
		for j := 0; j < len(num2)-1-i; j++ {
			cur += "0"
		}
		cur += singleMulti(num1, int(num2[i]-'0'))
		ans = addReverse(ans, cur)
	}
	result := ""
	for i := len(ans) - 1; i >= 0; i-- {
		result += string(ans[i])
	}
	return result
}

func addReverse(num1, num2 string) string {
	if len(num1) < len(num2) {
		return addReverse(num2, num1)
	}
	carry := 0
	ans := ""
	for i := 0; i < len(num1); i++ {
		cur := carry + int(num1[i]-'0')
		if i < len(num2) {
			cur += int(num2[i] - '0')
		}
		carry = cur / 10
		cur %= 10
		ans += string(byte(cur) + '0')
	}
	if carry != 0 {
		ans += string(byte(carry) + '0')
	}
	return ans
}

func singleMulti(num1 string, num2 int) string {
	fmt.Printf("multi %s * %d\n", num1, num2)
	carry := 0
	result := ""
	for i := len(num1) - 1; i >= 0; i-- {
		cur := int(num1[i]-'0')*num2 + carry
		carry = cur / 10
		result += string(byte(cur%10) + '0')
	}
	for carry > 0 {
		result += string(byte(carry%10) + '0')
		carry /= 10
	}
	return result
}

// @lc code=end
