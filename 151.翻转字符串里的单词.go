/*
 * @lc app=leetcode.cn id=151 lang=golang
 *
 * [151] 翻转字符串里的单词
 */
package main

// @lc code=start
func reverseWords(s string) string {
	l, r := 0, 0

	words := []string{}
	for r < len(s) {
		for r < len(s) && s[r] == ' ' {
			l++
			r++
		}
		if r == len(s) {
			break
		}
		for r < len(s) && s[r] != ' ' {
			r++
		}
		words = append(words, s[l:r])
		l = r
	}
	ans := ""
	for i := len(words) - 1; i >= 0; i-- {
		ans += words[i]
		if i != 0 {
			ans += " "
		}
	}
	return ans
}

// @lc code=end
