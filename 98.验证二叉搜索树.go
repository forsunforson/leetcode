/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
 */
package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	// nums := new([]int)
	// bstToString(root, nums)
	// for i := 0; i < len(*nums); i++ {
	// 	copyNums := *nums
	// 	if i != 0 && copyNums[i] <= copyNums[i-1] {
	// 		return false
	// 	}
	// 	if i != len(*nums)-1 && copyNums[i] >= copyNums[i+1] {
	// 		return false
	// 	}
	// }
	// return true
	return isValidBSTCore(root, math.MinInt64, math.MaxInt64)
}

func isValidBSTCore(root *TreeNode, min, max int64) bool {
	if root == nil {
		return true
	}
	if int64(root.Val) <= min || int64(root.Val) >= max {
		return false
	}
	return isValidBSTCore(root.Left, min, int64(root.Val)) && isValidBSTCore(root.Right, int64(root.Val), max)
}

func bstToString(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		bstToString(root.Left, nums)
	}
	*nums = append(*nums, root.Val)
	if root.Right != nil {
		bstToString(root.Right, nums)
	}
}

// @lc code=end
