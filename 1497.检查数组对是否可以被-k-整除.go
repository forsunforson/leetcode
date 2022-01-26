/*
 * @lc app=leetcode.cn id=1497 lang=golang
 *
 * [1497] 检查数组对是否可以被 k 整除
 */
package main

// @lc code=start
func canArrange(arr []int, k int) bool {
	// 检查余数是否可以凑对
	mod := make([]int, k)
	for _, v := range arr {
		rest := (v%k + k) % k
		mod[rest]++
	}
	for i := 0; i <= k/2; i++ {
		for mod[i] > 0 {
			mod[i]--
			j := (k - i) % k
			mod[j]--
		}
	}
	for _, m := range mod {
		if m != 0 {
			return false
		}
	}
	return true
}

// @lc code=end
