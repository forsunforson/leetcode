/*
 * @lc app=leetcode.cn id=1719 lang=golang
 *
 * [1719] 重构一棵树的方案数
 */
package main

import "math"

// @lc code=start
func checkWays(pairs [][]int) int {
	// union-find?
	// 构建一张图，如果存在一个节点是其他所有点的邻居，则返回1；存在多个节点，返回2；否则返回0
	// 实际上即使只存在一个根节点，子树可能能够替换，也存在返回2的可能
	mapping := map[int]map[int]bool{}
	for _, pair := range pairs {
		x, y := pair[0], pair[1]
		if mapping[x] == nil {
			mapping[x] = map[int]bool{}
		}
		mapping[x][y] = true
		if mapping[y] == nil {
			mapping[y] = map[int]bool{}
		}
		mapping[y][x] = true
	}
	root := -1
	// 一定存在一个点，是其他所有点的邻居
	for node, neightbors := range mapping {
		if len(neightbors) == len(mapping)-1 {
			root = node
			break
		}
	}
	if root == -1 {
		return 0
	}
	ans := 1
	for node, neighbours := range mapping {
		if node == root {
			continue
		}

		currDegree := len(neighbours)
		parent := -1
		parentDegree := math.MaxInt32
		// 根据 degree 的大小找到 node 的父节点 parent
		for neighbour := range neighbours {
			if len(mapping[neighbour]) < parentDegree && len(mapping[neighbour]) >= currDegree {
				// 找到degree比大于等于当前的最小邻居作为父节点
				parent = neighbour
				parentDegree = len(mapping[neighbour])
			}
		}
		if parent == -1 {
			return 0
		}
		// 检测 neighbours 是否为 adj[parent] 的子集
		for neighbour := range neighbours {
			if neighbour != parent && !mapping[parent][neighbour] {
				return 0
			}
		}
		// 如果degree相等，则存在多种方法
		if parentDegree == currDegree {
			ans = 2
		}
	}
	return ans

}

// @lc code=end
