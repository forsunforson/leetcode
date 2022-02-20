/*
 * @lc app=leetcode.cn id=785 lang=golang
 *
 * [785] 判断二分图
 */
package main

// @lc code=start
func isBipartite(graph [][]int) bool {
	// 对于每个节点，若它属于A，它的邻居都得是属于B
	// 0 未访问 1 A 2 B
	colors := map[int]int{}
	var dfs func(int, int)
	valid := true
	dfs = func(cur, color int) {
		colors[cur] = color
		neighborColor := 1
		if color == 1 {
			neighborColor = 2
		}
		for _, neighbor := range graph[cur] {
			if colors[neighbor] == 0 {
				dfs(neighbor, neighborColor)
				if !valid {
					return
				}
			} else if colors[neighbor] != neighborColor {
				valid = false
				return
			}
		}
	}
	for i := 0; i < len(graph); i++ {
		if colors[i] == 0 {
			dfs(i, 1)
		}
		if !valid {
			return false
		}
	}
	return true
}

// @lc code=end
