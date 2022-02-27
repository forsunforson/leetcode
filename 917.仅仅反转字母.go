/*
 * @lc app=leetcode.cn id=917 lang=golang
 *
 * [917] 仅仅反转字母
 */
package main

// @lc code=start
func reverseOnlyLetters(s string) string {
	b := []byte(s)
	l, r := 0, len(s)-1
	for l < r {
		if (b[l] < 'a' || b[l] > 'z') && (b[l] < 'A' || b[l] > 'Z') {
			l++
			continue
		}
		if (b[r] < 'a' || b[r] > 'z') && (b[r] < 'A' || b[r] > 'Z') {
			r--
			continue
		}
		b[l], b[r] = b[r], b[l]
		l++
		r--
	}
	return string(b)
}

// @lc code=end
