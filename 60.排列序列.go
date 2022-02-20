/*
 * @lc app=leetcode.cn id=60 lang=golang
 *
 * [60] 排列序列
 */
package main

import "strconv"

// @lc code=start
func getPermutation(n int, k int) string {
	ans := []string{}
	visited := make([]bool, n)
	var dfs func(string)
	dfs = func(cur string) {
		if len(cur) == n {
			ans = append(ans, cur)
			return
		}
		if len(ans) >= k {
			return
		}
		for i := 1; i <= n; i++ {
			if visited[i-1] {
				continue
			}
			visited[i-1] = true
			dfs(cur + strconv.Itoa(i))
			visited[i-1] = false
		}
	}
	dfs("")
	return ans[k-1]
}

// @lc code=end
