/*
 * @lc app=leetcode.cn id=1096 lang=golang
 *
 * [1096] 花括号展开 II
 */
package main

import (
	"sort"
	"strings"
)

// @lc code=start
func braceExpansionII(expression string) []string {
	ans := []string{}
	// BFS，从最内的一对括号开始匹配
	q := []string{expression}
	set := map[string]bool{}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if strings.Index(cur, "{") == -1 {
			set[cur] = true
			continue
		}
		l, r := 0, 0
		idx := 0
		for cur[idx] != '}' {
			if cur[idx] == '{' {
				l = idx
			}
			idx++
		}
		r = idx
		for _, str := range strings.Split(cur[l+1:r], ",") {
			builder := strings.Builder{}
			builder.WriteString(cur[:l])
			builder.WriteString(str)
			builder.WriteString(cur[r+1:])
			q = append(q, builder.String())
		}
	}
	for k := range set {
		ans = append(ans, k)
	}
	sort.Strings(ans)
	return ans
}

// @lc code=end
