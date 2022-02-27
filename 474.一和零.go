/*
 * @lc app=leetcode.cn id=474 lang=golang
 *
 * [474] 一和零
 */
package main

// @lc code=start
func findMaxForm(strs []string, m int, n int) int {
	// 背包问题，dp[i][j]代表当前包中有i个0 j个1，最大的子集数量
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	subset := 0
	// 在最外层循环strs，防止重复选择
	for _, str := range strs {
		zero, one := count(str)
		for i := m; i >= zero; i-- {
			for j := n; j >= one; j-- {
				dp[i][j] = max(dp[i][j], dp[i-zero][j-one]+1)
				subset = max(subset, dp[i][j])
			}
		}
	}
	return subset
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func count(str string) (int, int) {
	zero, one := 0, 0
	for i := range str {
		if str[i] == '0' {
			zero++
		} else {
			one++
		}
	}
	return zero, one
}

// @lc code=end
