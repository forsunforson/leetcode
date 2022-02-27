/*
 * @lc app=leetcode.cn id=187 lang=golang
 *
 * [187] 重复的DNA序列
 */
package main

// @lc code=start
func findRepeatedDnaSequences(s string) []string {
	if len(s) <= 10 {
		return []string{}
	}
	m := map[string]int{}
	l, r := 0, 10
	ans := []string{}
	for r <= len(s) {
		cur := s[l:r]
		m[cur]++
		if m[cur] == 2 {
			ans = append(ans, cur)
		}
		l++
		r++
	}
	return ans
}

// @lc code=end
