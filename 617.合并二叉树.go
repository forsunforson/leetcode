/*
 * @lc app=leetcode.cn id=617 lang=golang
 *
 * [617] 合并二叉树
 */
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	node := TreeNode{}
	if root1 == nil {
		node.Val = root2.Val
		node.Left = mergeTrees(nil, root2.Left)
		node.Right = mergeTrees(nil, root2.Right)
	} else if root2 == nil {
		node.Val = root1.Val
		node.Left = mergeTrees(root1.Left, nil)
		node.Right = mergeTrees(root1.Right, nil)
	} else {
		node.Val = root1.Val + root2.Val
		node.Left = mergeTrees(root1.Left, root2.Left)
		node.Right = mergeTrees(root1.Right, root2.Right)
	}

	return &node
}

// @lc code=end
