/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 */
package main

// @lc code=start
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for i := 0; i <= n; i++ {
		dp[0][i] = i
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// 对于左或上，因为新增了一个字符，操作肯定+1
			left := dp[i-1][j] + 1
			up := dp[i][j-1] + 1
			// 对于两个字符串都增加一个字符，如果相等，就不用操作
			leftUp := dp[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				leftUp += 1
			}
			dp[i][j] = min(leftUp, min(left, up))
		}
	}
	return dp[m][n]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end
