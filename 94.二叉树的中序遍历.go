/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
 */
package main

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	inorderCore(root, &ret)
	return ret
}

func inorderCore(root *TreeNode, nodeList *[]int) {
	if root == nil {
		return
	}
	inorderCore(root.Left, nodeList)
	(*nodeList) = append((*nodeList), root.Val)
	inorderCore(root.Right, nodeList)
}

// @lc code=end
