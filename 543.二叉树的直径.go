/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
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
func diameterOfBinaryTree(root *TreeNode) int {
	// 从题目描述看，像是左右子树最大深度相加
	ans := 1
	getMaxDepth(root, &ans)
	return ans - 1
}

func getMaxDepth(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}
	l := getMaxDepth(root.Left, ans)
	r := getMaxDepth(root.Right, ans)
	if *ans < l+r+1 {
		*ans = l + r + 1
	}
	ret := l
	if r > l {
		ret = r
	}
	return ret + 1
}

// @lc code=end
