/*
 * @lc app=leetcode.cn id=231 lang=golang
 *
 * [231] 2 的幂
 */
package main

// @lc code=start
func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	return n&-n == n
}

// @lc code=end
