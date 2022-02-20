/*
 * @lc app=leetcode.cn id=210 lang=golang
 *
 * [210] 课程表 II
 */
package main

// @lc code=start
func findOrder(numCourses int, prerequisites [][]int) []int {
	// 有向图，mapping[x][y] y表示先修课程
	mapping := make([][]int, numCourses)
	indegree := make([]int, numCourses)
	for i := range mapping {
		mapping[i] = make([]int, 0)
	}
	for _, pre := range prerequisites {
		mapping[pre[1]] = append(mapping[pre[1]], pre[0])
		indegree[pre[0]] += 1
	}
	// 对于有向图的BFS，每次放入入度为0的，相当于无依赖课程
	q := []int{}
	ans := []int{}
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		top := q[0]
		q = q[1:]
		ans = append(ans, top)
		for _, course := range mapping[top] {
			indegree[course] -= 1
			if indegree[course] == 0 {
				q = append(q, course)
			}
		}

	}
	if len(ans) != numCourses {
		return []int{}
	}
	return ans
}

// @lc code=end
