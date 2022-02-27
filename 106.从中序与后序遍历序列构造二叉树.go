/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 && len(postorder) == 0 {
		return nil
	}
	node := &TreeNode{
		Val: postorder[len(postorder)-1],
	}
	idx := 0
	for ; inorder[idx] != node.Val; idx++ {
	}
	node.Left = buildTree(inorder[:idx], postorder[:idx])

	if idx < len(inorder)-1 {
		node.Right = buildTree(inorder[idx+1:], postorder[idx:len(postorder)-1])
	}

	return node
}

// @lc code=end

