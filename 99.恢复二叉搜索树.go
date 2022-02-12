/*
 * @lc app=leetcode.cn id=99 lang=golang
 *
 * [99] 恢复二叉搜索树
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
func recoverTree(root *TreeNode) {
	// 空间复杂度O(n)的解法
	// 中序遍历，构建数组，然后交换逆序
	nodes := []*TreeNode{}
	addNode(root, &nodes)
	// 最差的方法，尝试交换每一对逆序的，再判断
	// for i := 0; i < len(nodes)-1; i++ {
	// 	for j := i + 1; j < len(nodes); j++ {
	// 		swap(nodes[i], nodes[j])
	// 		if isRight(nodes) {
	// 			return
	// 		}
	// 		swap(nodes[i], nodes[j])
	// 	}
	// }
	// 对于如果正好存在交换一对树就能恢复的解，那这两个数和前后存在逆序
	x, y := findReverse(nodes)
	swap(nodes[x], nodes[y])
}

func morrisVisit(root *TreeNode) {
	// 用于记录当前root的前驱节点
	var predecessor *TreeNode
	for root != nil {
		if root.Left != nil {
			predecessor = root.Left
			// 找到左子树最右节点
			for predecessor.Right != nil && predecessor.Right != root {
				predecessor = predecessor.Right
			}
			// 如果右子等于root，说明访问过
			if predecessor.Right == nil {
				predecessor.Right = root
				// 开始遍历左子树
				root = root.Left
			} else {
				predecessor.Right = nil
				root = root.Right
			}
		} else {
			// 如果左子为空，那么本身就是最后一个节点
			root = root.Right
		}
	}
}

func findReverse(nodes []*TreeNode) (int, int) {
	// x为第一个逆序，y为最后一个逆序
	x, y := -1, -1
	for i := 0; i < len(nodes)-1; i++ {
		if nodes[i].Val > nodes[i+1].Val {
			y = i + 1
			if x == -1 {
				x = i
			} else {
				break
			}
		}
	}
	return x, y
}

func addNode(root *TreeNode, nodes *[]*TreeNode) {
	if root == nil {
		return
	}
	addNode(root.Left, nodes)
	*nodes = append(*nodes, root)
	addNode(root.Right, nodes)
}

func swap(i, j *TreeNode) {
	i.Val, j.Val = j.Val, i.Val
}

func isRight(nodes []*TreeNode) bool {
	l, r := 0, 1
	for {
		if r >= len(nodes) {
			break
		}
		if nodes[l].Val >= nodes[r].Val {
			return false
		}
		l++
		r++
	}
	return true
}

// @lc code=end

