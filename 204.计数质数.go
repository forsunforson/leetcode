/*
 * @lc app=leetcode.cn id=204 lang=golang
 *
 * [204] 计数质数
 */
package main

// @lc code=start
func countPrimes(n int) int {
	if n == 0 || n == 1 {
		return 0
	}
	isPrime := make([]bool, n)
	count := 0
	for i := 0; i < n; i++ {
		isPrime[i] = true
	}
	for i := 2; i < n; i++ {
		if isPrime[i] {
			count++
			for j := 2; j*i < n; j++ {
				isPrime[i*j] = false
			}
		}
	}
	return count
}

// @lc code=end
