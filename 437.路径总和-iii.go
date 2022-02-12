/*
 * @lc app=leetcode.cn id=437 lang=golang
 *
 * [437] 路径总和 III
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

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	sum := rootSum(root, targetSum)
	sum += pathSum(root.Left, targetSum)
	sum += pathSum(root.Right, targetSum)
	return sum
}

// 以当前节点为根节点的路径
func rootSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	sum := 0
	if targetSum == root.Val {
		sum++
	}
	// 选择了根节点，必须连续选
	sum += rootSum(root.Left, targetSum-root.Val)
	sum += rootSum(root.Right, targetSum-root.Val)
	return sum
}

// @lc code=end

