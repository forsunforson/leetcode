/*
 * @lc app=leetcode.cn id=233 lang=golang
 *
 * [233] 数字 1 的个数
 */
package main

// @lc code=start
func countDigitOne(n int) int {
	// dp[i]表示i位数出现的1的总和
	ans := 0
	for k, mulk := 0, 1; n >= mulk; k++ {
		// k代表位数，mulk代表10^k
		ans += (n/(mulk*10))*mulk + min(max(n%(mulk*10)-mulk+1, 0), mulk)
		mulk *= 10
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
