package main

import "fmt"

// Input: edges = [[0,1],[1,2],[2,3],[1,4],[4,5]]
// Output: 4
// Explanation:
// A longest path of the tree is the path 3 - 2 - 1 - 4 - 5.
func Diameter(edges [][]int) int {
	// 以每个节点出发，进行深度优先遍历
	nodes := make([][]int, len(edges)+1)
	for i := range nodes {
		nodes[i] = make([]int, 0)
	}
	for _, edge := range edges {
		nodes[edge[0]] = append(nodes[edge[0]], edge[1])
		nodes[edge[1]] = append(nodes[edge[1]], edge[0])
	}
	visited := make([]bool, len(edges)+1)
	d := 0
	var dfs func(int, int)
	dfs = func(i, depth int) {
		if visited[i] {
			return
		}
		visited[i] = true
		depth += 1
		d = max(d, depth)
		for _, node := range nodes[i] {
			dfs(node, depth)
		}
		visited[i] = false
	}
	for i := 0; i <= len(edges); i++ {
		dfs(i, 0)
	}
	return d - 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	edges := [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 4}, {4, 5}}
	fmt.Println(Diameter(edges))
}
