/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
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

func maxPathSum(root *TreeNode) int {
	// 前缀树
	var MaxVal = math.MinInt32
	_ = buildTree(root, &MaxVal)
	return MaxVal
}

func buildTree(root *TreeNode, MaxVal *int) *TreeNode {
	if root == nil {
		return nil
	}
	node := TreeNode{}
	maxVal := root.Val
	l, r := 0, 0
	if root.Left != nil {
		node.Left = buildTree(root.Left, MaxVal)
		maxVal = max(maxVal, root.Val+node.Left.Val)
		l = max(l, node.Left.Val)
	}
	if root.Right != nil {
		node.Right = buildTree(root.Right, MaxVal)
		maxVal = max(maxVal, root.Val+node.Right.Val)
		r = max(r, node.Right.Val)
	}
	node.Val = maxVal
	*MaxVal = max(*MaxVal, root.Val+l+r)
	return &node
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

