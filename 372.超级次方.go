/*
 * @lc app=leetcode.cn id=372 lang=golang
 *
 * [372] 超级次方
 */
package main

// @lc code=start
func superPow(a int, b []int) int {
	if a == 1 {
		return 1
	}
	ans := 1
	for i := len(b) - 1; i >= 0; i-- {
		ans = ans * power(a, b[i]) % 1337
		// 倒序看，每次指数升一位，底数也得膨胀一位
		a = power(a, 10)
	}
	return ans
}

func power(x, i int) int {
	ans := 1
	for ; i > 0; i /= 2 {
		if i&1 == 1 {
			ans = x * ans % 1337
		}
		x = x * x % 1337
	}
	return ans
}

// @lc code=end
