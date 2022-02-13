/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	depth := math.MaxInt32
	if root.Left != nil {
		depth = min(depth, minDepth(root.Left))
	}
	if root.Right != nil {
		depth = min(depth, minDepth(root.Right))
	}
	return 1 + depth
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end

