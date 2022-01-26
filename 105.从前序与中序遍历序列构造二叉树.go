/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	// 先序：先本再左再右 中序：先左再本再右
	return buildTreeHelper(preorder, inorder)
}

func buildTreeHelper(pre, in []int) *TreeNode {
	if len(pre) == 0 || len(in) == 0 {
		return nil
	}
	if len(pre) == 1 && len(in) == 1 {
		return &TreeNode{
			Val: pre[0],
		}
	}
	root := TreeNode{
		Val: pre[0],
	}
	first := pre[0]
	idx := 0
	for {
		if first == in[idx] {
			break
		}
		idx++
	}
	root.Left = buildTreeHelper(pre[1:idx+1], in[0:idx])
	root.Right = buildTreeHelper(pre[idx+1:len(pre)], in[idx+1:len(in)])
	return &root
}

// @lc code=end
