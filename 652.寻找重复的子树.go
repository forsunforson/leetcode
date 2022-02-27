/*
 * @lc app=leetcode.cn id=652 lang=golang
 *
 * [652] 寻找重复的子树
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
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	// 序列化
	tree := map[string]int{}
	count := map[int]int{}
	ans := []*TreeNode{}
	t := 1
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		serial := fmt.Sprintf("%d_%d_%d", node.Val, dfs(node.Left), dfs(node.Right))
		uid := 0
		if v, ok := tree[serial]; !ok {
			tree[serial] = t
			uid = t
			t++
		} else {
			uid = v
		}

		count[uid]++
		if count[uid] == 2 {
			ans = append(ans, node)
		}
		return uid
	}
	dfs(root)
	return ans
}

// @lc code=end

