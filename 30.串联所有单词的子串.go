/*
 * @lc app=leetcode.cn id=30 lang=golang
 *
 * [30] 串联所有单词的子串
 */
package main

// @lc code=start
func findSubstring(s string, words []string) []int {
	// 滑动窗口匹配
	// words中的单词长度相同，所以滑动窗口的滑动step size可以保持word长度
	wordLen := len(words[0])
	wordLenSum := wordLen * len(words)
	dict := map[string]int{}
	for _, word := range words {
		dict[word]++
	}
	ans := []int{}
	l, r := 0, wordLenSum
	for r <= len(s) {
		curDict := map[string]int{}
		for i := 0; i < len(words); i++ {
			curWord := s[l+i*wordLen : l+(i+1)*wordLen]
			curDict[curWord]++
			if curDict[curWord] > dict[curWord] {
				break
			}
			if i == len(words)-1 {
				ans = append(ans, l)
			}
		}
		l += 1
		r += 1
	}
	return ans
}

// @lc code=end
