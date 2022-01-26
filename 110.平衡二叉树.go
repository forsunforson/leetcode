/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
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
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	curDepth := 0
	return isBalancedCore(root, &curDepth)
}

func isBalancedCore(root *TreeNode, depth *int) bool {
	leftDepth, rightDepth := (*depth)+1, (*depth)+1
	if root.Left != nil && !isBalancedCore(root.Left, &leftDepth) {
		return false
	}
	if root.Right != nil && !isBalancedCore(root.Right, &rightDepth) {
		return false
	}
	*depth = max(leftDepth, rightDepth)
	diff := abs(leftDepth - rightDepth)
	if diff > 1 {
		return false
	}
	return true
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func abs(i int) int {
	if i >= 0 {
		return i
	} else {
		return -i
	}
}

// @lc code=end

