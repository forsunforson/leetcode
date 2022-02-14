/*
 * @lc app=leetcode.cn id=567 lang=golang
 *
 * [567] 字符串的排列
 */
package main

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}
	l, r := 0, len(s1)-1
	char := make([]int, 26)
	dict := make([]int, 26)
	for i := 0; i < len(s1); i++ {
		char[s1[i]-'a']++
		dict[s2[i]-'a']++
	}
	if isEqual(char, dict) {
		return true
	}
	for {
		if r+1 == len(s2) {
			break
		}
		dict[s2[l]-'a']--
		dict[s2[r+1]-'a']++
		if isEqual(char, dict) {
			return true
		}
		l++
		r++
	}
	return false
}

func isEqual(a, b []int) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// @lc code=end
