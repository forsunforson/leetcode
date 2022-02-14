/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */
package main

// @lc code=start
func combine(n int, k int) [][]int {
	ans := [][]int{}
	var dfs func(int, int)
	path := []int{}
	dfs = func(idx, k int) {
		if k == 0 {
			ans = append(ans, append([]int{}, path...))
			return
		}
		if idx > n {
			return
		}
		dfs(idx+1, k)
		path = append(path, idx)
		dfs(idx+1, k-1)
		path = path[:len(path)-1]
	}
	dfs(1, k)
	return ans
}

// @lc code=end
