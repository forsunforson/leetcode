/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */
package main

// @lc code=start
func plusOne(digits []int) []int {
	var result []int
	carry := true
	for i := len(digits) - 1; i >= 0; i-- {
		if carry {
			digits[i]++
			carry = false
		}
		if digits[i] >= 10 {
			digits[i] -= 10
			carry = true
		}
		result = append(result, digits[i])
	}
	if carry {
		result = append(result, 1)
	}
	for i := 0; i < len(result)/2; i++ {
		tmp := result[i]
		result[i] = result[len(result)-i-1]
		result[len(result)-i-1] = tmp
	}
	return result
}

// @lc code=end
