/*
 * @lc app=leetcode.cn id=329 lang=golang
 *
 * [329] 矩阵中的最长递增路径
 */
package main

// @lc code=start
func longestIncreasingPath(matrix [][]int) int {
	// dfs 中寻找最长路径
	// 直接搜会超时，用带memo的搜索
	maxPath := 0
	m, n := len(matrix), len(matrix[0])
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
	}
	dirs := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	var dfs func(int, int) int
	dfs = func(x, y int) int {
		if memo[x][y] != 0 {
			return memo[x][y]
		}
		memo[x][y] = 1
		for _, dir := range dirs {
			newX, newY := x+dir[0], y+dir[1]
			if newX < 0 || newX >= m || newY < 0 || newY >= n || matrix[newX][newY] <= matrix[x][y] {
				continue
			}
			memo[x][y] = max(memo[x][y], dfs(newX, newY)+1)
		}
		return memo[x][y]
	}
	for i := range matrix {
		for j := range matrix[i] {
			maxPath = max(maxPath, dfs(i, j))
		}
	}
	return maxPath
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end
