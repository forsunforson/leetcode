/*
 * @lc app=leetcode.cn id=52 lang=golang
 *
 * [52] N皇后 II
 */
package main

// @lc code=start
func totalNQueens(n int) int {
	column := make(map[int]bool)
	diag1 := make(map[int]bool)
	diag2 := make(map[int]bool)
	count := 0
	var dfs func(int)
	dfs = func(x int) {
		if x == n {
			count++
			return
		}
		for i := 0; i < n; i++ {
			if column[i] || diag1[x+i] || diag2[x-i] {
				continue
			}
			column[i] = true
			diag1[x+i] = true
			diag2[x-i] = true
			dfs(x + 1)
			column[i] = false
			diag1[x+i] = false
			diag2[x-i] = false
		}
	}
	dfs(0)
	return count
}

// @lc code=end
