/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 */
package main

// @lc code=start
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	idx := -1
	l, r := 0, len(needle)
	for r <= len(haystack) {
		if haystack[l:r] == needle {
			idx = l
			break
		}
		l++
		r++
	}
	return idx
}

// @lc code=end
