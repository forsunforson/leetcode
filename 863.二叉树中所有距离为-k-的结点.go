/*
 * @lc app=leetcode.cn id=863 lang=golang
 *
 * [863] 二叉树中所有距离为 K 的结点
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
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	// 转换成图，将target的父节点变为子节点
	ans := []int{}
	parents := map[int]*TreeNode{}
	var findParent func(*TreeNode)
	findParent = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left != nil {
			parents[root.Left.Val] = root
			findParent(root.Left)
		}
		if root.Right != nil {
			parents[root.Right.Val] = root
			findParent(root.Right)
		}
	}
	findParent(root)
	var findAns func(*TreeNode, *TreeNode, int)
	findAns = func(root, from *TreeNode, dis int) {
		if root == nil {
			return
		}
		if dis == k {
			ans = append(ans, root.Val)
			return
		}
		if root.Left != nil && root.Left != from {
			findAns(root.Left, root, dis+1)
		}
		if root.Right != nil && root.Right != from {
			findAns(root.Right, root, dis+1)
		}
		if parents[root.Val] != nil && parents[root.Val] != from {
			findAns(parents[root.Val], root, dis+1)
		}
	}
	findAns(target, nil, 0)
	return ans
}

// @lc code=end

