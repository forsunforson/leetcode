/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 */
package main

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// 将一些点设置为0的dp
	dp := make([][]int, len(obstacleGrid))
	for i := range dp {
		dp[i] = make([]int, len(obstacleGrid[0]))
	}
	if obstacleGrid[0][0] == 0 {
		dp[0][0] = 1
	} else {
		return 0
	}
	for i := range dp {
		for j := range dp[0] {
			if obstacleGrid[i][j] == 1 {
				continue
			}
			if i > 0 {
				dp[i][j] += dp[i-1][j]
			}
			if j > 0 {
				dp[i][j] += dp[i][j-1]
			}
		}
		//fmt.Printf("%v\n", dp[i])
	}
	return dp[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}

// @lc code=end
