/*
 * @lc app=leetcode.cn id=132 lang=golang
 *
 * [132] 分割回文串 II
 */
package main

import (
	"math"
)

// @lc code=start
func minCut(s string) int {
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
		for j := range dp[i] {
			dp[i][j] = true
		}
	}
	// 利用dp获得任意s[i:j]是否为回文字符串
	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
		}
	}
	// 再次利用dp
	d := make([]int, len(s))
	for i := range d {
		if dp[0][i] {
			// 本身是回文，直接跳过
			continue
		}
		d[i] = math.MaxInt32
		for j := 0; j < i; j++ {
			if dp[j+1][i] && d[j]+1 < d[i] {
				d[i] = d[j] + 1
			}
		}
	}
	return d[len(s)-1]
}

// @lc code=end
