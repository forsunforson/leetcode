/*
 * @lc app=leetcode.cn id=509 lang=golang
 *
 * [509] 斐波那契数
 */
package main

// @lc code=start
func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	res := make([]int, n+1)
	res[0] = 0
	res[1] = 1
	for i := 2; i <= n; i++ {
		res[i] = res[i-1] + res[i-2]
	}
	return res[n]
}

// @lc code=end
