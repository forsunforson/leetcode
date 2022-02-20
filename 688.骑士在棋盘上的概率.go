/*
 * @lc app=leetcode.cn id=688 lang=golang
 *
 * [688] 骑士在棋盘上的概率
 */
package main

import "math"

// @lc code=start
func knightProbability(n int, k int, row int, column int) float64 {
	// dp 来确定有多少位置能够留在棋盘上,[step][x][y]
	dp := make([][][]float64, k+1)
	for step := range dp {
		dp[step] = make([][]float64, n)
		for x := 0; x < n; x++ {
			dp[step][x] = make([]float64, n)
			for y := 0; y < n; y++ {
				if step == 0 {
					dp[step][x][y] = 1
				} else {
					for _, pos := range nextPos(x, y) {
						if pos[0] >= 0 && pos[0] < n && pos[1] >= 0 && pos[1] < n {
							dp[step][x][y] += dp[step-1][pos[0]][pos[1]] / 8
						}
					}
				}
			}
		}
	}
	return dp[k][row][column]
}

// 回溯是没法过的
func dfs(n int, k int, row int, column int) float64 {
	// 1/8 ^ k * x 其中x是还在棋盘上的计数
	count := 0
	var dfs func(int, int, int)
	dfs = func(k, x, y int) {
		if k < 0 || x < 0 || x >= n || y < 0 || y >= n {
			return
		}
		if k == 0 {
			count++
			return
		}

		for _, p := range nextPos(x, y) {
			dfs(k-1, p[0], p[1])
		}
	}
	dfs(k, row, column)
	return math.Pow(0.125, float64(k)) * float64(count)
}

func nextPos(x, y int) [][]int {
	pos := [][]int{
		{x + 2, y + 1},
		{x + 2, y - 1},
		{x + 1, y + 2},
		{x + 1, y - 2},
		{x - 1, y + 2},
		{x - 1, y - 2},
		{x - 2, y + 1},
		{x - 2, y - 1},
	}
	return pos
}

// @lc code=end
