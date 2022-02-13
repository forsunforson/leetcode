/*
 * @lc app=leetcode.cn id=450 lang=golang
 *
 * [450] 删除二叉搜索树中的节点
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
func deleteNode(root *TreeNode, key int) *TreeNode {
	// 如果节点存在右子，与节点的右子树最左节点交换
	// 如果是叶节点直接删除
	// 如果存在左子，不存在右子，与左子树最右节点交换
	if root == nil {
		return nil
	}
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		// 当前节点需要删除
		if root.Left == nil && root.Right == nil {
			return nil
		}
		if root.Right != nil {
			root.Val = nextNode(root.Right).Val
			root.Right = deleteNode(root.Right, root.Val)
		} else {
			root.Val = preNode(root.Left).Val
			root.Left = deleteNode(root.Left, root.Val)
		}
	}
	return root
}

func nextNode(root *TreeNode) *TreeNode {
	for root.Left != nil {
		root = root.Left
	}
	return root
}

func preNode(root *TreeNode) *TreeNode {
	for root.Right != nil {
		root = root.Right
	}
	return root
}

// @lc code=end

