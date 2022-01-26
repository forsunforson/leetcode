/*
 * @lc app=leetcode.cn id=1358 lang=golang
 *
 * [1358] 包含所有三种字符的子字符串数目
 */
package main

// @lc code=start
func numberOfSubstrings(s string) int {
	// 双指针，滑动窗口
	cnt := []int{0, 0, 0}
	res := 0

	l, r := 0, 0
	for r < len(s) {
		cnt[s[r]-'a']++
		for l < r && cnt[0] > 0 && cnt[1] > 0 && cnt[2] > 0 {
			res += len(s) - r
			cnt[s[l]-'a']--
			l++
		}
		r++
	}
	return res
}

// @lc code=end
