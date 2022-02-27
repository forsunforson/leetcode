/*
 * @lc app=leetcode.cn id=623 lang=golang
 *
 * [623] 在二叉树中增加一行
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
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if root == nil {
		return nil
	}
	if depth == 1 {
		node := &TreeNode{
			Val:  val,
			Left: root,
		}
		return node
	}
	if depth == 2 {
		newLeft := &TreeNode{
			Val: val,
		}
		if root.Left != nil {
			newLeft.Left = root.Left
		}
		root.Left = newLeft
		newRight := &TreeNode{Val: val}
		if root.Right != nil {
			newRight.Right = root.Right
		}
		root.Right = newRight
	}
	addOneRow(root.Left, val, depth-1)
	addOneRow(root.Right, val, depth-1)
	return root
}

// @lc code=end

