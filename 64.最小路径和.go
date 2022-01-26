/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 */
package main

import "math"

// @lc code=start
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			right, down := math.MaxInt32, math.MaxInt32
			if i+1 < m {
				right = dp[i+1][j]
			}
			if j+1 < n {
				down = dp[i][j+1]
			}
			if i == m-1 && j == n-1 {
				dp[i][j] = grid[i][j]
				continue
			}
			dp[i][j] = grid[i][j] + min(right, down)
		}
	}
	return dp[0][0]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end
