/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */
package main

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	l := 0
	count := 0
	words := map[byte]int{}
	for r := 1; r <= len(s); r++ {
		words[s[r-1]]++
		if words[s[r-1]] == 2 {
			for words[s[r-1]] == 2 {
				words[s[l]]--
				l++
			}
		}
		count = max(count, r-l)
	}
	return count
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// @lc code=end
