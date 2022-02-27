/*
 * @lc app=leetcode.cn id=201 lang=golang
 *
 * [201] 数字范围按位与
 */
package main

// @lc code=start
func rangeBitwiseAnd(left int, right int) int {
	// 最长公共前缀
	for right > left {
		right = right & (right - 1)
	}
	return right
}

// @lc code=end
