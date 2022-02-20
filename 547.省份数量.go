/*
 * @lc app=leetcode.cn id=547 lang=golang
 *
 * [547] 省份数量
 */
package main

// @lc code=start
// union-find
type node struct {
	name   int
	father *node
}

func union(a, b *node) {
	for a != a.father {
		a = a.father
	}
	for b != b.father {
		b = b.father
	}
	b.father = a
}
func isUnion(a, b *node) bool {
	for a != a.father {
		a = a.father
	}
	for b != b.father {
		b = b.father
	}
	return a == b
}
func getFather(a *node) *node {
	for a != a.father {
		a = a.father
	}
	return a
}
func unionfind(isConnected [][]int) int {
	nodes := make([]*node, len(isConnected))
	for i := range nodes {
		n := &node{
			name: i,
		}
		n.father = n
		nodes[i] = n
	}
	for i := range isConnected {
		for j := range isConnected[i] {
			if isConnected[i][j] == 1 {
				union(nodes[i], nodes[j])
			}
		}
	}
	father := make([]bool, len(isConnected))
	count := 0
	for i := range isConnected {
		name := getFather(nodes[i]).name
		if !father[name] {
			count++
			father[name] = true
		}
	}
	return count
}

func findCircleNum(isConnected [][]int) int {
	visited := make([]bool, len(isConnected))
	count := 0
	var tmp func(int)
	tmp = func(cur int) {
		if visited[cur] {
			return
		}
		visited[cur] = true
		for i := range isConnected[cur] {
			if isConnected[cur][i] == 1 {
				tmp(i)
			}
		}
	}
	for i := range isConnected {
		if !visited[i] {
			tmp(i)
			count++
		}
	}
	return count
}

// @lc code=end
