/*
 * @lc app=leetcode.cn id=174 lang=golang
 *
 * [174] 地下城游戏
 */
package main

import (
	"math"
)

// @lc code=start
func calculateMinimumHP(dungeon [][]int) int {
	// dp 最佳路径上 最小值
	m, n := len(dungeon), len(dungeon[0])
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[m][n-1], dp[m-1][n] = 1, 1
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			// 从需要血比较少的位置中选出一条，但是血最少也得是1
			minn := min(dp[i+1][j], dp[i][j+1])
			dp[i][j] = max(minn-dungeon[i][j], 1)
		}
	}
	return dp[0][0]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end
