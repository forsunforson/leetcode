/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 */
package main

// @lc code=start
func myPow(x float64, n int) float64 {
	// 根据n的奇、偶性进行递归
	if n == 0 {
		return 1
	}
	if n < 0 {
		return myPow(1/x, -n)
	}
	if n == 1 {
		return x
	}
	if n%2 == 1 {
		return x * myPow(x, n-1)
	} else {
		return myPow(x*x, n/2)
	}
}

// @lc code=end
