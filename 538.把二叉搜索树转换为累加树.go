/*
 * @lc app=leetcode.cn id=538 lang=golang
 *
 * [538] 把二叉搜索树转换为累加树
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
func convertBST(root *TreeNode) *TreeNode {
	// 先右子 再本结点 再左子
	return convertBST2(root)
}

func convertBST2(root *TreeNode) *TreeNode {
	sum := 0
	convertBST2Core(root, &sum)
	return root
}

func convertBST2Core(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	if root.Right != nil {
		convertBST2Core(root.Right, sum)
	}
	*sum = root.Val + *sum
	root.Val = *sum
	if root.Left != nil {
		convertBST2Core(root.Left, sum)
	}
}

// 最傻的方法，用List
func convertBST1(root *TreeNode) *TreeNode {
	nodeList := make([]*TreeNode, 0)
	buildList(&nodeList, root)
	sum := 0
	for i := len(nodeList) - 1; i >= 0; i-- {
		sum += nodeList[i].Val
		nodeList[i].Val = sum
	}
	return root
}

func buildList(nodeList *[]*TreeNode, root *TreeNode) {
	if root == nil {
		return
	}
	buildList(nodeList, root.Left)
	(*nodeList) = append((*nodeList), root)
	buildList(nodeList, root.Right)
}

// @lc code=end
