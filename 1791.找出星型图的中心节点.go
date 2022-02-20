/*
 * @lc app=leetcode.cn id=1791 lang=golang
 *
 * [1791] 找出星型图的中心节点
 */
package main

// @lc code=start
func findCenter(edges [][]int) int {
	degree := make([]int, len(edges)+2)
	for _, edge := range edges {
		degree[edge[0]]++
		degree[edge[1]]++
	}
	for i, d := range degree {
		if d == len(edges) {
			return i
		}
	}
	return 0
}

// @lc code=end
