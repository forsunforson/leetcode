/*
 * @lc app=leetcode.cn id=226 lang=golang
 *
 * [226] 翻转二叉树
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
func invertTree(root *TreeNode) *TreeNode {
	invertTreeCore(root)
	return root
}

func invertTreeCore(node *TreeNode) {
	if node == nil {
		return
	}
	tmpNode := node.Left
	node.Left = node.Right
	node.Right = tmpNode
	invertTreeCore(node.Left)
	invertTreeCore(node.Right)
}

// @lc code=end
