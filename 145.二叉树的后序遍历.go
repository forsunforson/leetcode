/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
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
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ans := []int{}
	var helper func(*TreeNode)
	helper = func(n *TreeNode) {
		if n == nil {
			return
		}
		helper(n.Left)
		helper(n.Right)
		ans = append(ans, n.Val)
	}
	helper(root)
	return ans
}

// @lc code=end

