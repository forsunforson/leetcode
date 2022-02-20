/*
 * @lc app=leetcode.cn id=129 lang=golang
 *
 * [129] 求根节点到叶节点数字之和
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	results := [][]int{}
	var dfs func(*TreeNode, []int)
	dfs = func(root *TreeNode, path []int) {
		if root == nil {
			return
		}
		path = append(path, root.Val)
		if root.Left == nil && root.Right == nil {
			results = append(results, append([]int{}, path...))
			return
		}
		if root.Left != nil {
			dfs(root.Left, path)
		}
		if root.Right != nil {
			dfs(root.Right, path)
		}
	}
	dfs(root, []int{})
	sum := 0
	for _, result := range results {
		tmp := 0
		for _, i := range result {
			tmp = 10*tmp + i
		}
		sum += tmp
	}
	return sum
}

// @lc code=end

