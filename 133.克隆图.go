/*
 * @lc app=leetcode.cn id=133 lang=golang
 *
 * [133] 克隆图
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	nodes := make(map[int]*Node)
	return helper(node, nodes)
}
func helper(node *Node, nodes map[int]*Node) *Node {
	if nodes[node.Val] != nil {
		return nodes[node.Val]
	}
	copyNode := &Node{
		Val:       node.Val,
		Neighbors: []*Node{},
	}
	nodes[node.Val] = copyNode
	for _, neighbor := range node.Neighbors {
		copyNode.Neighbors = append(copyNode.Neighbors, helper(neighbor, nodes))
	}
	return copyNode
}

// @lc code=end

