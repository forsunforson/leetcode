/*
 * @lc app=leetcode.cn id=990 lang=golang
 *
 * [990] 等式方程的可满足性
 */
package main

// @lc code=start
type node struct {
	val    string
	father *node
}

func isUnion(n1, n2 *node) bool {
	for n1 != n1.father {
		n1 = n1.father
	}
	for n2 != n2.father {
		n2 = n2.father
	}
	return n1 == n2
}
func union(n1, n2 *node) {
	for n2 != n2.father {
		father := n2.father
		n2.father = father.father
		n2 = father
	}
	n2.father = n1
}
func equationsPossible(equations []string) bool {
	// 实现union find
	nodes := map[byte]*node{}
	for _, equation := range equations {
		n1, ok := nodes[equation[0]]
		if !ok {
			n1 = &node{val: string(equation[0])}
			n1.father = n1
			nodes[equation[0]] = n1
		}
		n2, ok := nodes[equation[3]]
		if !ok {
			n2 = &node{val: string(equation[3])}
			n2.father = n2
			nodes[equation[3]] = n2
		}
		if equation[1] == '=' {
			if !isUnion(n1, n2) {
				union(n1, n2)
			}

		}

	}
	for _, equation := range equations {
		if equation[1] == '=' {
			continue
		}
		n1, n2 := nodes[equation[0]], nodes[equation[3]]
		if isUnion(n1, n2) {
			return false
		}
	}
	return true
}

// @lc code=end
