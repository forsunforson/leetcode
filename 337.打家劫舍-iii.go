/*
 * @lc app=leetcode.cn id=337 lang=golang
 *
 * [337] 打家劫舍 III
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
func rob(root *TreeNode) int {
	// 回溯
	return max(robHelper(root, true), robHelper(root, false))
}

func robHelper(root *TreeNode, rob bool) int {
	if root == nil {
		return 0
	}
	ret := root.Val
	if rob {
		return ret + robHelper(root.Left, false) + robHelper(root.Right, false)
	} else {
		left := max(robHelper(root.Left, true), robHelper(root.Left, false))
		right := max(robHelper(root.Right, true), robHelper(root.Right, false))
		return left + right
	}
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func rob(root *TreeNode) int {
	// 回溯
	ret := robHelper(root)
	return max(ret[0], ret[1])
}

func robHelper(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}
	// 减少两次访问
	l, r := robHelper(root.Left), robHelper(root.Right)
	ret := root.Val
	selected := ret + l[1] + r[1]
	unseleccted := max(l[0], l[1]) + max(r[0], r[1])
	return []int{selected, unseleccted}
}

// @lc code=end
