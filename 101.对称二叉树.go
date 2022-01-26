/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
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
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isBalancedCore(root.Left, root.Right)
}
func isBalancedCore(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left != nil && right == nil {
		return false
	}
	if left == nil && right != nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return isBalancedCore(left.Left, right.Right) && isBalancedCore(left.Right, right.Left)
}

// @lc code=end

