/*
 * @lc app=leetcode.cn id=172 lang=golang
 *
 * [172] 阶乘后的零
 */
package main

// @lc code=start
func trailingZeroes(n int) int {
	count := 0
	for i := 1; i <= n; i++ {
		count += countFive(i)
	}
	return count
}

func countFive(n int) int {
	count := 0
	for n > 0 {
		if n%5 == 0 {
			count++
			n /= 5
		} else {
			break
		}
	}
	return count
}

// @lc code=end
