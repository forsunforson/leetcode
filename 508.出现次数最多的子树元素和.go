/*
 * @lc app=leetcode.cn id=508 lang=golang
 *
 * [508] 出现次数最多的子树元素和
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
func findFrequentTreeSum(root *TreeNode) []int {
	// 后序遍历
	if root == nil {
		return []int{}
	}
	result := make(map[int]int)
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node.Left == nil && node.Right == nil {
			result[node.Val]++
			return node.Val
		}
		l, r := 0, 0
		if node.Left != nil {
			l = dfs(node.Left)
		}
		if node.Right != nil {
			r = dfs(node.Right)
		}
		result[node.Val+l+r]++
		return node.Val + l + r
	}
	dfs(root)
	ans := []int{}
	maxCount := 0
	for _, v := range result {
		maxCount = max(maxCount, v)
	}
	for k, v := range result {
		if v == maxCount {
			ans = append(ans, k)
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

