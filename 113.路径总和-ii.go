/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
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
func pathSum(root *TreeNode, targetSum int) [][]int {
	ans := make([][]int, 0)
	pathSumHelper(&ans, []int{}, root, targetSum)
	return ans
}

func pathSumHelper(ans *[][]int, path []int, root *TreeNode, targetSum int) {
	if root == nil {
		return
	}
	path = append(path, root.Val)
	if root.Left == nil && root.Right == nil && targetSum == root.Val {
		copyPath := append([]int{}, path...)
		*ans = append(*ans, copyPath)
		return
	}

	pathSumHelper(ans, path, root.Left, targetSum-root.Val)
	pathSumHelper(ans, path, root.Right, targetSum-root.Val)
}

// @lc code=end

