/*
 * @lc app=leetcode.cn id=222 lang=golang
 *
 * [222] 完全二叉树的节点个数
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
func countNodes(root *TreeNode) int {
	// 对于满二叉树，计算方法为2^h-1
	// 完全二叉树，需要遍历
	// 最左路径和最右路径长度是否相同来判断是否为满二叉树
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	if isFullTree(root) {
		depth := getDepth(root)
		count := 1<<depth - 1
		return count
	} else {
		return 1 + countNodes(root.Left) + countNodes(root.Right)
	}
}

func getDepth(root *TreeNode) int {
	count := 0
	for root != nil {
		count++
		root = root.Left
	}
	return count
}

func isFullTree(root *TreeNode) bool {
	return false
}

// @lc code=end

