/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
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
func flatten(root *TreeNode) {
	list := make([]*TreeNode, 0)
	visitNode(root, &list)
	for i, node := range list {
		node.Left = nil

		if i == len(list)-1 {
			node.Right = nil
		} else {
			node.Right = list[i+1]
		}
	}
}

func visitNode(root *TreeNode, list *[]*TreeNode) {
	if root == nil {
		return
	}
	*list = append(*list, root)
	visitNode(root.Left, list)
	visitNode(root.Right, list)
}

// @lc code=end

