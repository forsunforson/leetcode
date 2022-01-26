/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */
package main

// @lc code=start
func longestPalindrome(s string) string {
	if len(s) < 1 {
		return ""
	}
	maxLength := 0
	var palindrome string
	for idx := range s {
		r1, l1 := expand(s, idx, idx)
		r2, l2 := expand(s, idx, idx+1)
		var r, l int
		if r1-l1 > r2-l2 {
			r, l = r1, l1
		} else {
			r, l = r2, l2
		}
		if maxLength < r-l+1 {
			palindrome = s[l : r+1]
			maxLength = r - l + 1
		}
	}
	return palindrome
}

// expand 根据中心点位置，返回构成回文串的start，end
func expand(s string, left, right int) (int, int) {
	for left >= 0 && right <= len(s)-1 && s[left] == s[right] {
		left--
		right++

	}
	return right - 1, left + 1
}

// @lc code=end
